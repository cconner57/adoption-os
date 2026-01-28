package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Invitation struct {
	Token     string    `json:"token"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type InvitationModel struct {
	DB *sql.DB
}

func (m InvitationModel) Insert(invite *Invitation) error {
	query := `
		INSERT INTO invitations (token, email, role, expires_at)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at`

	args := []any{invite.Token, invite.Email, invite.Role, invite.ExpiresAt}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&invite.CreatedAt)
}

func (m InvitationModel) Get(token string) (*Invitation, error) {
	query := `
		SELECT token, email, role, expires_at, created_at
		FROM invitations
		WHERE token = $1`

	var invite Invitation

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, token).Scan(
		&invite.Token,
		&invite.Email,
		&invite.Role,
		&invite.ExpiresAt,
		&invite.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &invite, nil
}

func (m InvitationModel) Delete(token string) error {
	query := `
		DELETE FROM invitations
		WHERE token = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, token)
	return err
}
