package data

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"time"
)

type Session struct {
	Token     string
	UserID    string
	Expiry    time.Time
	IP        string
	UserAgent string
}

type SessionModel struct {
	DB *sql.DB
}

func (m SessionModel) Insert(userID string, ttl time.Duration, ip, userAgent string) (string, error) {
	token, err := generateToken()
	if err != nil {
		return "", err
	}

	expiry := time.Now().Add(ttl)

	query := `
		INSERT INTO sessions (token, user_id, expiry, ip, user_agent)
		VALUES ($1, $2, $3, $4, $5)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = m.DB.ExecContext(ctx, query, token, userID, expiry, ip, userAgent)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m SessionModel) Get(token string) (*Session, error) {
	query := `
		SELECT token, user_id, expiry, ip, user_agent
		FROM sessions
		WHERE token = $1 AND expiry > CURRENT_TIMESTAMP`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var s Session

	err := m.DB.QueryRowContext(ctx, query, token).Scan(
		&s.Token,
		&s.UserID,
		&s.Expiry,
		&s.IP,
		&s.UserAgent,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if no valid session found
		}
		return nil, err
	}

	return &s, nil
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
