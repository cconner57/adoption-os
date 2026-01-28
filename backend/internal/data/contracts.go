package data

import (
	"context"
	"database/sql"
	"time"
)

type Contract struct {
	Token         string     `json:"token"`
	ApplicationID int64      `json:"application_id"`
	Type          string     `json:"type"`
	Signature     *string    `json:"signature"`
	SignedAt      *time.Time `json:"signed_at"`
	ExpiresAt     time.Time  `json:"expires_at"`
	CreatedAt     time.Time  `json:"created_at"`
}

type ContractModel struct {
	DB *sql.DB
}

func (m ContractModel) Insert(contract *Contract) error {
	query := `
		INSERT INTO contracts (token, application_id, type, expires_at)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at`

	args := []any{
		contract.Token,
		contract.ApplicationID,
		contract.Type,
		contract.ExpiresAt,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&contract.CreatedAt)
}

func (m ContractModel) Get(token string) (*Contract, error) {
	query := `
		SELECT token, application_id, type, expires_at, created_at, signature, signed_at
		FROM contracts
		WHERE token = $1`

	var contract Contract

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, token).Scan(
		&contract.Token,
		&contract.ApplicationID,
		&contract.Type,
		&contract.ExpiresAt,
		&contract.CreatedAt,
		&contract.Signature,
		&contract.SignedAt,
	)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &contract, nil
}

func (m ContractModel) GetByApplicationID(appID int64) (*Contract, error) {
	query := `
		SELECT token, application_id, type, expires_at, created_at, signature, signed_at
		FROM contracts
		WHERE application_id = $1
		ORDER BY created_at DESC
		LIMIT 1`

	var contract Contract

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, appID).Scan(
		&contract.Token,
		&contract.ApplicationID,
		&contract.Type,
		&contract.ExpiresAt,
		&contract.CreatedAt,
		&contract.Signature,
		&contract.SignedAt,
	)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &contract, nil
}

func (m ContractModel) UpdateSignature(token string, signature string) error {
	query := `
		UPDATE contracts 
		SET signature = $1, signed_at = NOW()
		WHERE token = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, signature, token)
	return err
}
