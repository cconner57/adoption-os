package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

	"github.com/google/uuid"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type User struct {
	ID           string
	CreatedAt    time.Time
	Name         string
	Email        string
	PasswordHash string // Argon2 hash
	Activated    bool
	Version      int
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(user *User) error {
	query := `
		INSERT INTO users (id, name, email, password_hash, activated)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version`

	// Generate UUID if not present? Or always?
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	args := []any{user.ID, user.Name, user.Email, user.PasswordHash, user.Activated}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
	if err != nil {
		switch {
		// case err.Error() contains duplicate key?
		// We should have specific check, but for now generic return
		default:
			return err
		}
	}

	return nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `
		SELECT id, created_at, name, email, password_hash, activated, version
		FROM users
		WHERE email = $1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Activated,
		&user.Version,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (m UserModel) Get(id string) (*User, error) {
	query := `
		SELECT id, created_at, name, email, password_hash, activated, version
		FROM users
		WHERE id = $1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Activated,
		&user.Version,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &user, nil
}
