package data

import "database/sql"

type Models struct {
	Volunteers   VolunteerModel
	Users        UserModel
	Pets         PetModel
	Metrics      MetricModel
	Sessions     SessionModel
	Shifts       ShiftModel
	Applications ApplicationModel
	Marketing    MarketingModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Volunteers:   VolunteerModel{DB: db},
		Users:        UserModel{DB: db},
		Pets:         PetModel{DB: db},
		Metrics:      MetricModel{DB: db},
		Sessions:     SessionModel{DB: db},
		Shifts:       ShiftModel{DB: db},
		Applications: ApplicationModel{DB: db},
		Marketing:    MarketingModel{DB: db},
	}
}
