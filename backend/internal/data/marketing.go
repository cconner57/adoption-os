package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type Participant struct {
	ID            string `json:"id"`
	TicketNumber  int    `json:"ticketNumber"`
	Name          string `json:"name"`
	Contact       string `json:"contact"`
	Paid          bool   `json:"paid"`
	Date          string `json:"date,omitempty"`
	PaymentMethod string `json:"paymentMethod,omitempty"`
	Amount        string `json:"amount,omitempty"`
	Comments      string `json:"comments,omitempty"`
}

type Campaign struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Status      string          `json:"status"`
	StartDate   string          `json:"startDate"`
	EndDate     string          `json:"endDate"`
	Goal        string          `json:"goal"`
	Progress    int             `json:"progress"`
	Metric      string          `json:"metric"`
	Type        string          `json:"type"`
	Prize       string          `json:"prize,omitempty"`
	TicketPrice float64         `json:"ticketPrice,omitempty"`
	WinnerID    string          `json:"winnerId,omitempty"`
	Participants json.RawMessage `json:"participants,omitempty"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

type MarketingModel struct {
	DB *sql.DB
}

func (m MarketingModel) GetCampaigns(status string) ([]*Campaign, error) {
	if m.DB == nil {
		return nil, fmt.Errorf("database connection not available")
	}

	query := `
		SELECT 
			id, name, status, start_date, end_date, goal, progress, metric, type, 
			prize, ticket_price, winner_id, COALESCE(participants, '[]'), created_at, updated_at
		FROM marketing_campaigns
		WHERE 1=1
	`
	args := []interface{}{}
	if status != "" {
		query += " AND status = $1"
		args = append(args, status)
	}

	query += " ORDER BY created_at DESC"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	campaigns := []*Campaign{}
	for rows.Next() {
		var c Campaign
		var prize, winnerId sql.NullString
		var ticketPrice sql.NullFloat64

		err := rows.Scan(
			&c.ID, &c.Name, &c.Status, &c.StartDate, &c.EndDate, &c.Goal, &c.Progress, &c.Metric, &c.Type,
			&prize, &ticketPrice, &winnerId, &c.Participants, &c.CreatedAt, &c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if prize.Valid {
			c.Prize = prize.String
		}
		if ticketPrice.Valid {
			c.TicketPrice = ticketPrice.Float64
		}
		if winnerId.Valid {
			c.WinnerID = winnerId.String
		}

		campaigns = append(campaigns, &c)
	}

	return campaigns, nil
}

func (m MarketingModel) GetCampaign(id string) (*Campaign, error) {
	if m.DB == nil {
		return nil, fmt.Errorf("database connection not available")
	}

	query := `
		SELECT 
			id, name, status, start_date, end_date, goal, progress, metric, type, 
			prize, ticket_price, winner_id, COALESCE(participants, '[]'), created_at, updated_at
		FROM marketing_campaigns
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var c Campaign
	var prize, winnerId sql.NullString
	var ticketPrice sql.NullFloat64

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&c.ID, &c.Name, &c.Status, &c.StartDate, &c.EndDate, &c.Goal, &c.Progress, &c.Metric, &c.Type,
		&prize, &ticketPrice, &winnerId, &c.Participants, &c.CreatedAt, &c.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	if prize.Valid {
		c.Prize = prize.String
	}
	if ticketPrice.Valid {
		c.TicketPrice = ticketPrice.Float64
	}
	if winnerId.Valid {
		c.WinnerID = winnerId.String
	}

	return &c, nil
}

func (m MarketingModel) CreateCampaign(c *Campaign) error {
	if m.DB == nil {
		return fmt.Errorf("database connection not available")
	}

	query := `
		INSERT INTO marketing_campaigns (
			id, name, status, start_date, end_date, goal, progress, metric, type, 
			prize, ticket_price, winner_id, participants, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, NOW(), NOW()
		)
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query,
		c.ID, c.Name, c.Status, c.StartDate, c.EndDate, c.Goal, c.Progress, c.Metric, c.Type,
		c.Prize, c.TicketPrice, c.WinnerID, c.Participants,
	)
	return err
}

func (m MarketingModel) UpdateCampaign(c *Campaign) error {
	if m.DB == nil {
		return fmt.Errorf("database connection not available")
	}

	query := `
		UPDATE marketing_campaigns 
		SET 
			name = $1, status = $2, start_date = $3, end_date = $4, goal = $5, 
			progress = $6, metric = $7, type = $8, prize = $9, ticket_price = $10, 
			winner_id = $11, participants = $12, updated_at = NOW()
		WHERE id = $13
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query,
		c.Name, c.Status, c.StartDate, c.EndDate, c.Goal, c.Progress, c.Metric, c.Type,
		c.Prize, c.TicketPrice, c.WinnerID, c.Participants, c.ID,
	)
	return err
}
