package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

type Metric struct {
	ID        int64           `json:"id"`
	EventType string          `json:"eventType"`
	EventData json.RawMessage `json:"eventData"`
	CreatedAt time.Time       `json:"createdAt"`
}

type MetricModel struct {
	DB *sql.DB
}

func (m MetricModel) Insert(metric *Metric) error {
	query := `
		INSERT INTO metrics (event_type, event_data)
		VALUES ($1, $2)
		RETURNING id, created_at`

	args := []any{metric.EventType, metric.EventData}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&metric.ID, &metric.CreatedAt)
}
