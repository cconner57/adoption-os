package main

import (
	"sort"
	"strings"
	"time"
)

func (app *application) recalculateVolunteerStats(volunteerID int64) error {
	shifts, err := app.models.Shifts.GetForVolunteer(volunteerID)
	if err != nil {
		return err
	}

	currentYear := time.Now().Year()

	// Filter shifts for current year only
	var currentYearShifts []struct {
		Timestamp time.Time
		Status    string
		StartTime string
		EndTime   string
		Notes     string
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

		// Strictly filter for current year
		if date.Year() != currentYear {
			continue
		}

		// Parse Time (Attempt 24h then 12h)
		t, err := time.Parse("15:04", s.StartTime)
		if err != nil {
			t, err = time.Parse("3:04 PM", s.StartTime)
		}
		// If time fails, default to midnight
		if err != nil {
			t = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
		}

		// Combine Date + Time
		timestamp := time.Date(
			date.Year(), date.Month(), date.Day(),
			t.Hour(), t.Minute(), 0, 0,
			time.Local,
		)

		currentYearShifts = append(currentYearShifts, struct {
			Timestamp time.Time
			Status    string
			StartTime string
			EndTime   string
			Notes     string
		}{
			Timestamp: timestamp,
			Status:    s.Status,
			StartTime: s.StartTime,
			EndTime:   s.EndTime,
			Notes:     s.Notes,
		})
	}

	// 1. Calculate Total Hours (Current Year)
	var totalHours float64
	for _, s := range currentYearShifts {
		if s.Status == "completed" || s.Status == "all_good" || s.Status == "late" {
			start, err1 := time.Parse("15:04", s.StartTime)
			end, err2 := time.Parse("15:04", s.EndTime)

			if err1 != nil {
				start, err1 = time.Parse("3:04 PM", s.StartTime)
			}
			if err2 != nil {
				end, err2 = time.Parse("3:04 PM", s.EndTime)
			}

			if err1 == nil && err2 == nil {
				duration := end.Sub(start).Hours()
				if duration < 0 {
					duration += 24
				}
				if duration > 0 {
					totalHours += duration
				}
			}
		}
	}

	// 2. Calculate Reliability (Current Year)
	score := 0.0
	// Base points
	// Completed: +20
	// Late: +10
	// Missed: -50
	// Covered: -5 / -10 / -20
	// Bonus: +10 / +20

	for _, s := range currentYearShifts {
		// Base Score
		switch s.Status {
		case "completed", "all_good":
			score += 20
		case "late":
			score += 10
		case "missed", "no_show":
			score -= 50
		case "covered", "covered_24h", "covered 24h":
			score -= 5
		case "covered_less_24h", "covered_late", "covered <24h notice", "covered late":
			score -= 10
		case "covered_less_1h", "covered <1h notice":
			score -= 20
		}

		// Bonus Points from Notes via "Covering for"
		// Logic matches src/utils/reliability.ts
		if strings.Contains(s.Notes, "Covering for") {
			if strings.Contains(s.Notes, ">24h notice") {
				score += 10
			} else if strings.Contains(s.Notes, "<24h notice") {
				score += 20
			}
		}
	}

	// 3. Calculate Streak (Current Year)
	// Sort descending for streak calc
	sort.Slice(currentYearShifts, func(i, j int) bool {
		return currentYearShifts[i].Timestamp.After(currentYearShifts[j].Timestamp)
	})

	streak := 0
	for _, s := range currentYearShifts {
		if s.Status == "completed" || s.Status == "all_good" || s.Status == "late" {
			streak++
		} else {
			break
		}
	}

	// Cap at 100 on the upper end (allow negatives? Frontend doesn't explicitly clamp negatives but display might. We'll leave as is)
	if score > 100 {
		score = 100
	}
	reliability := int(score)

	return app.models.Volunteers.UpdateStats(volunteerID, reliability, int(totalHours), streak)
}
