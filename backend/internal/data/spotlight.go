package data

import (
	"context"
	"fmt"
	"time"
)

func (m PetModel) validateSpotlight(p *Pet, isSpotlight bool) error {
	if isSpotlight {
		// Count OTHER pets that are featured AND available
		countQuery := `SELECT count(*) FROM pets WHERE profile_settings->>'isSpotlightFeatured' = 'true' AND id != $1 AND status = 'available'`
		var count int
		ctxCount, cancelCount := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelCount()

		err := m.DB.QueryRowContext(ctxCount, countQuery, p.ID).Scan(&count)
		if err != nil {
			return err
		}

		if count >= MaxSpotlightPets {
			return fmt.Errorf("maximum of %d spotlight pets allowed. Please unfeature another pet first", MaxSpotlightPets)
		}
	}
	return nil
}
