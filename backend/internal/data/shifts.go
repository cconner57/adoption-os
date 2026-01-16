package data

import (
	"context"
	"database/sql"
	"time"
)

type Shift struct {
	ID          int64  `json:"id"`
	VolunteerID int64  `json:"volunteerId"`
	Date        string `json:"date"` // Keeping as string YYYY-MM-DD for simplicity as requested
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Role        string `json:"role"`
	Status      string `json:"status"`
	Notes       string `json:"notes"`
	Version     int    `json:"-"` // Not using version yet but good practice
}

type ShiftModel struct {
	DB *sql.DB
}

func (m ShiftModel) Insert(shift *Shift) error {
	query := `
		INSERT INTO shifts (volunteer_id, date, start_time, end_time, role, status, notes)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, version`

	args := []any{
		shift.VolunteerID,
		shift.Date,
		shift.StartTime,
		shift.EndTime,
		shift.Role,
		shift.Status,
		shift.Notes,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&shift.ID, new(time.Time), &shift.Version)
}

func (m ShiftModel) GetForVolunteer(volunteerID int64) ([]*Shift, error) {
	query := `
		SELECT id, volunteer_id, TO_CHAR(date, 'YYYY-MM-DD'), start_time, end_time, role, status, notes
		FROM shifts
		WHERE volunteer_id = $1
		ORDER BY date DESC, start_time ASC`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, volunteerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []*Shift
	for rows.Next() {
		var s Shift
		err := rows.Scan(
			&s.ID,
			&s.VolunteerID,
			&s.Date,
			&s.StartTime,
			&s.EndTime,
			&s.Role,
			&s.Status,
			&s.Notes,
		)
		if err != nil {
			return nil, err
		}
		shifts = append(shifts, &s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return shifts, nil
}

func (m ShiftModel) Get(id int64) (*Shift, error) {
	query := `
		SELECT id, volunteer_id, TO_CHAR(date, 'YYYY-MM-DD'), start_time, end_time, role, status, notes, version
		FROM shifts
		WHERE id = $1`

	var shift Shift

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&shift.ID,
		&shift.VolunteerID,
		&shift.Date,
		&shift.StartTime,
		&shift.EndTime,
		&shift.Role,
		&shift.Status,
		&shift.Notes,
		&shift.Version,
	)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &shift, nil
}

func (m ShiftModel) Update(shift *Shift) error {
	query := `
		UPDATE shifts
		SET date = $1, start_time = $2, end_time = $3, role = $4, status = $5, notes = $6, updated_at = NOW(), version = version + 1
		WHERE id = $7 AND version = $8
		RETURNING version`

	args := []any{
		shift.Date,
		shift.StartTime,
		shift.EndTime,
		shift.Role,
		shift.Status,
		shift.Notes,
		shift.ID,
		shift.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&shift.Version)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m ShiftModel) Delete(id int64) error {
	query := `
		DELETE FROM shifts
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m ShiftModel) GetRoleCounts() (map[string]int, error) {
	query := `
		SELECT role, COUNT(*) as count
		FROM shifts
		GROUP BY role`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make(map[string]int)
	for rows.Next() {
		var role string
		var count int
		if err := rows.Scan(&role, &count); err != nil {
			return nil, err
		}
		counts[role] = count
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return counts, nil
}
