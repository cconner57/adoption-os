package data

import (
	"context"
	"database/sql"
	"time"
)

type NotificationSubscription struct {
	ID        int64     `json:"id"`
	Endpoint  string    `json:"endpoint"`
	P256dh    string    `json:"keys_p256dh"`
	Auth      string    `json:"keys_auth"`
	UserAgent string    `json:"user_agent,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type NotificationModel struct {
	DB *sql.DB
}

// Insert adds a new subscription to the database.
// It uses an UPSERT (ON CONFLICT) strategy on endpoint to avoid duplicates but update keys if they changed.
func (m NotificationModel) Insert(sub *NotificationSubscription) error {
	query := `
		INSERT INTO notification_subscriptions (endpoint, p256dh, auth, user_agent, created_at)
		VALUES ($1, $2, $3, $4, NOW())
		ON CONFLICT (endpoint) DO UPDATE
		SET p256dh = EXCLUDED.p256dh,
			auth = EXCLUDED.auth,
			user_agent = EXCLUDED.user_agent,
			created_at = NOW()`

	args := []interface{}{
		sub.Endpoint,
		sub.P256dh,
		sub.Auth,
		sub.UserAgent,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}

// Delete removes a subscription by endpoint (e.g., if it expired or user unsubscribed).
func (m NotificationModel) Delete(endpoint string) error {
	query := `
		DELETE FROM notification_subscriptions
		WHERE endpoint = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, endpoint)
	return err
}

// GetAll retrieves all active subscriptions.
func (m NotificationModel) GetAll() ([]*NotificationSubscription, error) {
	query := `
		SELECT id, endpoint, p256dh, auth, user_agent, created_at
		FROM notification_subscriptions`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []*NotificationSubscription
	for rows.Next() {
		var sub NotificationSubscription
		err := rows.Scan(
			&sub.ID,
			&sub.Endpoint,
			&sub.P256dh,
			&sub.Auth,
			&sub.UserAgent,
			&sub.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		subs = append(subs, &sub)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subs, nil
}
