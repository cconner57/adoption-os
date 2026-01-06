package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/validator"
	"github.com/lib/pq"
)

type VolunteerApplication struct {
	ID                    int64     `json:"id"`
	CreatedAt             time.Time `json:"createdAt"`
	FirstName             string    `json:"firstName"`
	LastName              string    `json:"lastName"`
	Address               string    `json:"address"`
	City                  string    `json:"city"`
	Zip                   string    `json:"zip"`
	PhoneNumber           string    `json:"phoneNumber"`
	Birthday              string    `json:"birthday"`
	Age                   *int      `json:"age"`
	Allergies             bool      `json:"allergies"`
	EmergencyContactName  string    `json:"emergencyContactName"`
	EmergencyContactPhone string    `json:"emergencyContactPhone"`
	VolunteerExperience   string    `json:"volunteerExperience"`
	InterestReason        string    `json:"interestReason"`
	PositionPreferences   []string  `json:"positionPreferences"`
	Availability          []string  `json:"availability"`
	NameFull              string    `json:"nameFull"`
	SignatureData         *string   `json:"signatureData"`
	SignatureDate         string    `json:"signatureDate"`
	ParentName            string    `json:"parentName"`
	ParentSignatureData   *string   `json:"parentSignatureData"`
	ParentSignatureDate   string    `json:"parentSignatureDate"`
	Status                string    `json:"status"`
}

type VolunteerModel struct {
	DB *sql.DB
}

func (m VolunteerModel) Insert(application *VolunteerApplication) error {
	query := `
		INSERT INTO volunteer_applications (
			first_name, last_name, address, city, zip, phone_number, birthday, age,
			allergies, emergency_contact_name, emergency_contact_phone,
			volunteer_experience, interest_reason, position_preferences, availability,
			full_name_signature, signature_data, signature_date,
			parent_name, parent_signature_data, parent_signature_date
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
		RETURNING id, created_at, status`

	args := []any{
		application.FirstName, application.LastName, application.Address, application.City, application.Zip, application.PhoneNumber, application.Birthday, application.Age,
		application.Allergies, application.EmergencyContactName, application.EmergencyContactPhone,
		application.VolunteerExperience, application.InterestReason, pq.Array(application.PositionPreferences), pq.Array(application.Availability),
		application.NameFull, application.SignatureData, application.SignatureDate,
		application.ParentName, application.ParentSignatureData, application.ParentSignatureDate,
	}

	return m.DB.QueryRowContext(context.Background(), query, args...).Scan(&application.ID, &application.CreatedAt, &application.Status)
}

func ValidateVolunteerApplication(v *validator.Validator, application *VolunteerApplication) {
	v.Check(application.FirstName != "", "firstName", "must be provided")
	v.Check(application.LastName != "", "lastName", "must be provided")
	v.Check(application.Address != "", "address", "must be provided")
	v.Check(application.City != "", "city", "must be provided")
	v.Check(application.Zip != "", "zip", "must be provided")
	v.Check(application.PhoneNumber != "", "phoneNumber", "must be provided")
	v.Check(application.Birthday != "", "birthday", "must be provided")

	var isUnder21 bool
	if application.Birthday != "" {
		t, err := time.Parse("2006-01-02", application.Birthday)
		if err != nil {
			// Try parsing MM/DD/YYYY
			t, err = time.Parse("01/02/2006", application.Birthday)
		}

		if err == nil {
			now := time.Now()
			age := now.Year() - t.Year()
			if now.YearDay() < t.YearDay() {
				age--
			}
			if age < 21 {
				isUnder21 = true
			}

			if application.Age == nil {
				ageInt := age
				application.Age = &ageInt
			}
		}
	}

	if isUnder21 {
		v.Check(application.Age != nil, "age", "must be provided for applicants under 21")
		if application.Age != nil {
			v.Check(*application.Age >= 0, "age", "must be a positive number")
		}
	}

	v.Check(application.EmergencyContactName != "", "emergencyContactName", "must be provided")
	v.Check(application.EmergencyContactPhone != "", "emergencyContactPhone", "must be provided")
	v.Check(application.InterestReason != "", "interestReason", "must be provided")
	v.Check(len(application.PositionPreferences) > 0, "positionPreferences", "must select at least one position")
	v.Check(len(application.Availability) > 0, "availability", "must select at least one availability slot")
	v.Check(application.NameFull != "", "nameFull", "must be provided")
	v.Check(application.SignatureDate != "", "signatureDate", "must be provided")
	v.Check(application.SignatureData != nil && *application.SignatureData != "", "signatureData", "must be provided")
	if isUnder21 {
		v.Check(application.ParentName != "", "parentName", "must be provided for applicants under 21")
		v.Check(application.ParentSignatureDate != "", "parentSignatureDate", "must be provided for applicants under 21")
		v.Check(application.ParentSignatureData != nil && *application.ParentSignatureData != "", "parentSignatureData", "must be provided for applicants under 21")
	}
}
