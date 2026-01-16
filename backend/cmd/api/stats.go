package main

import (
	"sort"
	"time"
)

func (app *application) recalculateVolunteerStats(volunteerID int64) error {
	shifts, err := app.models.Shifts.GetForVolunteer(volunteerID)
	if err != nil {
		return err
	}

	// 1. Calculate Total Hours
	var totalHours float64
	for _, s := range shifts {
		if s.Status == "completed" || s.Status == "all_good" || s.Status == "late" {
			// Try parsing 24h first (most common in this app)
			start, err1 := time.Parse("15:04", s.StartTime)
			end, err2 := time.Parse("15:04", s.EndTime)

			// Fallback to 12h if needed
			if err1 != nil {
				start, err1 = time.Parse("3:04 PM", s.StartTime)
			}
			if err2 != nil {
				end, err2 = time.Parse("3:04 PM", s.EndTime)
			}

			if err1 == nil && err2 == nil {
				duration := end.Sub(start).Hours()
				// Handle overnight shifts if negative? Assuming single day for now as per requirement complexity.
				if duration < 0 {
					duration += 24
				}
				if duration > 0 {
					totalHours += duration
				}
			}
		}
	}

	// 2. Calculate Reliability
	// Base: 0
	// Completed: +20
	// Late: +10
	// Covered >24h: -5
	// Covered <24h: -10
	// Covered <1h: -20
	// Missed: -50
	// Cap at 100. Allow negatives.

	score := 0.0
	// 3. Calculate Streak
	var pastShifts []struct {
		Timestamp time.Time
		Status    string
	}

	for _, s := range shifts {
		// Skip future/scheduled shifts
		if s.Status == "scheduled" {
			continue
		}

		// Parse Date
		date, err := time.Parse("2006-01-02", s.Date)
		if err != nil {
			continue
		}

		// Parse Time (Attempt 24h then 12h)
		t, err := time.Parse("15:04", s.StartTime)
		if err != nil {
			t, err = time.Parse("3:04 PM", s.StartTime)
		}
		// If time fails, default to midnight (00:00) so it at least works
		if err != nil {
			t = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
		}

		// Combine Date + Time
		timestamp := time.Date(
			date.Year(), date.Month(), date.Day(),
			t.Hour(), t.Minute(), 0, 0,
			time.Local, // Use Local to align with user expectation
		)

		pastShifts = append(pastShifts, struct {
			Timestamp time.Time
			Status    string
		}{Timestamp: timestamp, Status: s.Status})

		// Calculate Reliability Score
		switch s.Status {
		case "completed", "all_good":
			score += 20
		case "late":
			score += 10
		case "missed", "no_show":
			score -= 50
		case "covered", "covered_24h", "covered 24h":
			score -= 5
		case "covered_less_24h", "covered_late", "covered <24h notice":
			score -= 10
		case "covered_less_1h", "covered <1h notice":
			score -= 20
		}
	}

	// Sort descending (Most recent first)
	sort.Slice(pastShifts, func(i, j int) bool {
		return pastShifts[i].Timestamp.After(pastShifts[j].Timestamp)
	})

	streak := 0
	if len(pastShifts) > 0 {
		for _, s := range pastShifts {
			if s.Status == "completed" || s.Status == "all_good" {
				streak++
			} else {
				// Any other status (late, missed, covered) breaks the streak
				break
			}
		}
	}

	if score > 100 {
		score = 100
	}
	reliability := int(score)

	return app.models.Volunteers.UpdateStats(volunteerID, reliability, int(totalHours), streak)
}
