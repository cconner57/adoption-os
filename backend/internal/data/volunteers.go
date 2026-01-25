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
	Email                 string    `json:"email"`
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
	// Honeypot
	FaxNumber string `json:"fax_number"`
}

type VolunteerModel struct {
	DB *sql.DB
}

func (m VolunteerModel) Insert(application *VolunteerApplication) error {
	query := `
		INSERT INTO volunteer_applications (
			first_name, last_name, email, address, city, zip, phone_number, birthday, age,
			allergies, emergency_contact_name, emergency_contact_phone,
			volunteer_experience, interest_reason, position_preferences, availability,
			full_name_signature, signature_data, signature_date,
			parent_name, parent_signature_data, parent_signature_date
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
		RETURNING id, created_at, status`

	args := []any{
		application.FirstName, application.LastName, application.Email, application.Address, application.City, application.Zip, application.PhoneNumber, application.Birthday, application.Age,
		application.Allergies, application.EmergencyContactName, application.EmergencyContactPhone,
		application.VolunteerExperience, application.InterestReason, pq.Array(application.PositionPreferences), pq.Array(application.Availability),
		application.NameFull, application.SignatureData, application.SignatureDate,
		application.ParentName, application.ParentSignatureData, application.ParentSignatureDate,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&application.ID, &application.CreatedAt, &application.Status)
}

func ValidateVolunteerApplication(v *validator.Validator, application *VolunteerApplication) {
	v.Check(application.FirstName != "", "firstName", "must be provided")
	v.Check(application.LastName != "", "lastName", "must be provided")
	v.Check(application.Email != "", "email", "must be provided")
	if application.Email != "" {
		v.Check(validator.Matches(application.Email, validator.EmailRX), "email", "must be a valid email address")
	}
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

// --- Active Volunteer Management ---

type Volunteer struct {
	ID                    int64     `json:"id"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	FirstName             string    `json:"firstName"`
	LastName              string    `json:"lastName"`
	Email                 string    `json:"email"`
	Phone                 string    `json:"phone"`
	Address               string    `json:"address"`
	City                  string    `json:"city"`
	Zip                   string    `json:"zip"`
	Role                  string    `json:"role"`
	Status                string    `json:"status"`
	Bio                   string    `json:"bio"`
	PhotoURL              string    `json:"photoUrl"`
	ReliabilityScore      int       `json:"reliabilityScore"`
	TotalHours            int       `json:"totalHours"`
	Streak                int       `json:"streak"`
	JoinDate              string    `json:"joinDate"` // Stored as date, simplified to string for API
	Allergies             bool      `json:"allergies"`
	Skills                []string  `json:"skills"`
	PositionPreferences   []string  `json:"positionPreferences"`
	Availability          []string  `json:"availability"`
	Badges                []string  `json:"badges"`
	Birthday              string    `json:"birthday"`
	EmergencyContactName  string    `json:"emergencyContactName"`
	EmergencyContactPhone string    `json:"emergencyContactPhone"`
	InterestReason        string    `json:"interestReason"`
	VolunteerExperience   string    `json:"volunteerExperience"`
	Version               int       `json:"version"`
}

func (m VolunteerModel) InsertGetId(v *Volunteer) error {
	query := `
		INSERT INTO volunteers (
			first_name, last_name, email, phone, address, city, zip, role, status,
			bio, photo_url, reliability_score, total_hours, streak, join_date, allergies,
			skills, position_preferences, availability, badges,
			birthday, emergency_contact_name, emergency_contact_phone, interest_reason, volunteer_experience
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)
		RETURNING id, created_at, updated_at, version`

	// Helper to handle empty strings for dates
	var joinDate any = v.JoinDate
	if v.JoinDate == "" {
		joinDate = nil
	}
	var birthday any = v.Birthday
	if v.Birthday == "" {
		birthday = nil
	}

	args := []any{
		v.FirstName, v.LastName, v.Email, v.Phone, v.Address, v.City, v.Zip, v.Role, v.Status,
		v.Bio, v.PhotoURL, v.ReliabilityScore, v.TotalHours, v.Streak, joinDate, v.Allergies,
		pq.Array(v.Skills), pq.Array(v.PositionPreferences), pq.Array(v.Availability), pq.Array(v.Badges),
		birthday, v.EmergencyContactName, v.EmergencyContactPhone, v.InterestReason, v.VolunteerExperience,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&v.ID, &v.CreatedAt, &v.UpdatedAt, &v.Version)
}

func (m VolunteerModel) GetAll(firstName string, lastName string, role string, status string, filters Filters) ([]*Volunteer, Metadata, error) {
	// Updated query to include filters
	query := `
		SELECT count(*) OVER(), id, created_at, updated_at, first_name, last_name, email, phone, address, city, zip, role, status,
		bio, photo_url, reliability_score, total_hours, streak, join_date, allergies, skills, position_preferences, availability, badges,
		COALESCE(TO_CHAR(birthday, 'YYYY-MM-DD'), '') as birthday,
		COALESCE(emergency_contact_name, '') as emergency_contact_name,
		COALESCE(emergency_contact_phone, '') as emergency_contact_phone,
		COALESCE(interest_reason, '') as interest_reason,
		COALESCE(volunteer_experience, '') as volunteer_experience,
		version
		FROM volunteers
		WHERE (to_tsvector('simple', first_name) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (LOWER(role) = LOWER($2) OR $2 = '')
		AND (LOWER(status) = LOWER($3) OR $3 = '')
		ORDER BY id DESC`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{firstName, role, status}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	volunteers := []*Volunteer{}

	for rows.Next() {
		var v Volunteer
		var joinDate time.Time // Scan as time, format later

		err := rows.Scan(
			&totalRecords,
			&v.ID,
			&v.CreatedAt,
			&v.UpdatedAt,
			&v.FirstName,
			&v.LastName,
			&v.Email,
			&v.Phone,
			&v.Address,
			&v.City,
			&v.Zip,
			&v.Role,
			&v.Status,
			&v.Bio,
			&v.PhotoURL,
			&v.ReliabilityScore,
			&v.TotalHours,
			&v.Streak,
			&joinDate,
			&v.Allergies,
			pq.Array(&v.Skills),
			pq.Array(&v.PositionPreferences),
			pq.Array(&v.Availability),
			pq.Array(&v.Badges),
			&v.Birthday,
			&v.EmergencyContactName,
			&v.EmergencyContactPhone,
			&v.InterestReason,
			&v.VolunteerExperience,
			&v.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		v.JoinDate = joinDate.Format("2006-01-02")
		volunteers = append(volunteers, &v)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return volunteers, metadata, nil
}

func (m VolunteerModel) Get(id int64) (*Volunteer, error) {
	query := `
		SELECT id, created_at, updated_at, first_name, last_name, email, phone, address, city, zip, role, status,
		bio, photo_url, reliability_score, total_hours, streak, join_date, allergies, skills, position_preferences, availability, badges,
		COALESCE(TO_CHAR(birthday, 'YYYY-MM-DD'), '') as birthday, emergency_contact_name, emergency_contact_phone, interest_reason, volunteer_experience, version
		FROM volunteers
		WHERE id = $1`

	var v Volunteer
	var joinDate time.Time

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&v.ID, &v.CreatedAt, &v.UpdatedAt, &v.FirstName, &v.LastName, &v.Email, &v.Phone,
		&v.Address, &v.City, &v.Zip, &v.Role, &v.Status,
		&v.Bio, &v.PhotoURL, &v.ReliabilityScore, &v.TotalHours, &v.Streak,
		&joinDate, &v.Allergies, pq.Array(&v.Skills), pq.Array(&v.PositionPreferences), pq.Array(&v.Availability), pq.Array(&v.Badges),
		&v.Birthday, &v.EmergencyContactName, &v.EmergencyContactPhone, &v.InterestReason, &v.VolunteerExperience,
		&v.Version,
	)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	v.JoinDate = joinDate.Format("2006-01-02")
	return &v, nil
}

func (m VolunteerModel) Update(v *Volunteer) error {
	query := `
		UPDATE volunteers
		SET first_name = $1, last_name = $2, email = $3, phone = $4, address = $5, city = $6, zip = $7,
		    role = $8, status = $9, bio = $10, photo_url = $11, reliability_score = $12, total_hours = $13,
		    streak = $14, join_date = $15, allergies = $16, skills = $17, position_preferences = $18,
		    availability = $19, badges = $20, birthday = $21, emergency_contact_name = $22,
			emergency_contact_phone = $23, interest_reason = $24, volunteer_experience = $25,
			updated_at = NOW(), version = version + 1
		WHERE id = $26 AND version = $27
		RETURNING updated_at, version`

	// Helper to handle empty strings for dates
	var joinDate any = v.JoinDate
	if v.JoinDate == "" {
		joinDate = nil
	}
	var birthday any = v.Birthday
	if v.Birthday == "" {
		birthday = nil
	}

	args := []any{
		v.FirstName, v.LastName, v.Email, v.Phone, v.Address, v.City, v.Zip,
		v.Role, v.Status, v.Bio, v.PhotoURL, v.ReliabilityScore, v.TotalHours,
		v.Streak, joinDate, v.Allergies, pq.Array(v.Skills), pq.Array(v.PositionPreferences),
		pq.Array(v.Availability), pq.Array(v.Badges),
		birthday, v.EmergencyContactName, v.EmergencyContactPhone, v.InterestReason, v.VolunteerExperience,
		v.ID, v.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&v.UpdatedAt, &v.Version)
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

func (m VolunteerModel) UpdateStats(id int64, reliability int, hours int, streak int) error {
	query := `
		UPDATE volunteers
		SET reliability_score = $1, total_hours = $2, streak = $3, updated_at = NOW(), version = version + 1
		WHERE id = $4`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, reliability, hours, streak, id)
	return err
}
