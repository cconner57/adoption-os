package data

import "database/sql"

type Models struct {
	Volunteers VolunteerModel
	Pets       PetModel
	Metrics    MetricModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Volunteers: VolunteerModel{DB: db},
		Pets:       PetModel{DB: db},
		Metrics:    MetricModel{DB: db},
	}
}
