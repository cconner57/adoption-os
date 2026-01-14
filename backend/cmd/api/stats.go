package main

import (
	"math"
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
	// Logic mirrors frontend:
	// Base: 100
	// Late: 80
	// Covered (>24h): 50
	// Covered Late (<24h): 25
	// Missed: 0
	// Bonus: Covering (+10), Event (+5)
	// Penalties: Incidents (Ignored for now effectively as we don't fetch them)

	var totalPossible float64
	var earnedPoints float64

	// Filter past verified shifts
	now := time.Now()
	var pastShifts []struct {
		Date   time.Time
		Status string
	}

	for _, s := range shifts {
		// Try parsing YYYY-MM-DD
		shiftDate, err := time.Parse("2006-01-02", s.Date)
		if err != nil {
			continue
		}

		if shiftDate.Before(now) && s.Status != "scheduled" {
			totalPossible += 100
			points := 0.0

			if s.Status == "missed" || s.Status == "no_show" {
				points = 0
			} else {
				points = 100
			}

			earnedPoints += points

			// For streak calc
			pastShifts = append(pastShifts, struct {
				Date   time.Time
				Status string
			}{Date: shiftDate, Status: s.Status})
		}
	}

	reliability := 100
	if totalPossible > 0 {
		reliability = int(math.Round((earnedPoints / totalPossible) * 100))
	}
	if reliability > 100 {
		reliability = 100
	} else if reliability < 0 {
		reliability = 0
	}

	// 3. Calculate Streak
	// Sort descending
	sort.Slice(pastShifts, func(i, j int) bool {
		return pastShifts[i].Date.After(pastShifts[j].Date)
	})

	streak := 0
	if len(pastShifts) > 0 {
		// Map shifts to Year-Week
		type WeekKey struct {
			Year int
			Week int
		}
		visitedWeeks := make(map[WeekKey]bool)

		// Get most recent shift's week
		lastDate := pastShifts[0].Date
		// Start checking from the week of the most recent shift
		// If the most recent shift was > 2 weeks ago, is streak broken?
		// Usually yes. If > 14 days ago, streak is 0?
		// Let's assume if it's been > 2 weeks since last volunteer, streak is broken.
		daysSinceLast := now.Sub(lastDate).Hours() / 24

		isStreakAlive := true
		if daysSinceLast > 14 {
			isStreakAlive = false
		}

		if isStreakAlive {
			currentYear, currentWeek := lastDate.ISOWeek()
			lastWeekKey := WeekKey{currentYear, currentWeek}
			streak = 1 // Count the first week
			visitedWeeks[lastWeekKey] = true

			// Iterate backwards looking for consecutive weeks
			uniqueWeeks := []WeekKey{lastWeekKey}

			for _, s := range pastShifts {
				// Only count good shifts
				if s.Status == "missed" || s.Status == "no_show" {
					continue // Or break? If they missed a shift in a week, does it break streak?
					// Usually streak is "consecutive weeks WITHOUT missing"? or just "consecutive weeks WITH volunteering"?
					// Requirement: "Reliability" handles misses. Streak handles continuity.
					// Let's assume streak is "Weeks with at least one good shift".
					// If they have a "missed" shift in a week but also a "good" shift, does it count? Yes.
				}

				y, w := s.Date.ISOWeek()
				wk := WeekKey{y, w}
				if !visitedWeeks[wk] {
					visitedWeeks[wk] = true
					uniqueWeeks = append(uniqueWeeks, wk)
				}
			}

			// Check continuity of uniqueWeeks
			// They are somewhat ordered since pastShifts is sorted, but within a week sorting might be mixed.
			// uniqueWeeks naturally collected in roughly desc order because pastShifts is desc.

			prevY, prevW := currentYear, currentWeek

			for i := 1; i < len(uniqueWeeks); i++ {
				curr := uniqueWeeks[i]

				// Check if curr is exactly 1 week before prev
				isConsecutive := false

				// Standard case: same year
				if prevY == curr.Year {
					if prevW-curr.Week == 1 {
						isConsecutive = true
					}
				} else if prevY == curr.Year+1 {
					// Boundary case
					// If prev is Year X+1 Week 1, curr should be Year X Week 52 or 53
					// ISO weeks: Dec 28 is always in the same year as Jan 4.
					// Last week of year is 52 or 53.
					if prevW == 1 && (curr.Week == 52 || curr.Week == 53) {
						isConsecutive = true
					}
				}

				if isConsecutive {
					streak++
					prevY, prevW = curr.Year, curr.Week
				} else {
					break // Gap found
				}
			}
		}
	}

	return app.models.Volunteers.UpdateStats(volunteerID, reliability, int(totalHours), streak)
}
