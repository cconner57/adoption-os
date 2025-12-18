package models

import (
	"time"
)

// Pet matches the TypeScript IPet interface exactly
type Pet struct {
	ID         string    `json:"id"`
	Slug       *string   `json:"slug"` // Optional in TS
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Name       string    `json:"name"`
	Species    string    `json:"species"`
	Sex        string    `json:"sex"`
	LitterName *string   `json:"litterName"` // Optional in TS

	// Nested JSON Structures
	Physical     Physical     `json:"physical"`
	Behavior     Behavior     `json:"behavior"`
	Medical      Medical      `json:"medical"`
	Descriptions Descriptions `json:"descriptions"`
	Details      Details      `json:"details"`
	Adoption     Adoption     `json:"adoption"`
	Foster       Foster       `json:"foster"`
	Returned     Returned     `json:"returned"`
	Sponsored    Sponsored    `json:"sponsored"`
	Photos       []Photo      `json:"photos"`

	// Profile settings (Admin only usually, but part of your interface)
	ProfileSettings ProfileSettings `json:"profileSettings"`
}

// --- Nested Structs ---

type Physical struct {
	AgeGroup      *string  `json:"ageGroup"`
	Breed         *string  `json:"breed"`
	CoatLength    *string  `json:"coatLength"`
	Color         *string  `json:"color"`
	DateOfBirth   *string  `json:"dateOfBirth"`
	Size          *string  `json:"size"`
	CurrentWeight *float64 `json:"currentWeight"`
}

type Behavior struct {
	Bonded               *BondedInfo `json:"bonded"` // Optional object
	EnergyLevel          *string     `json:"energyLevel"`
	HealthSummary        *string     `json:"healthSummary"`
	IsGoodWithCats       *bool       `json:"isGoodWithCats"`
	IsGoodWithDogs       *bool       `json:"isGoodWithDogs"`
	IsGoodWithKids       *bool       `json:"isGoodWithKids"`
	IsHouseTrained       *bool       `json:"isHouseTrained"`
	MustGoWithAnotherCat *bool       `json:"mustGoWithAnotherCat"`
	MustGoWithAnotherDog *bool       `json:"mustGoWithAnotherDog"`
	PersonalityTags      []string    `json:"personalityTags"`
	PrefersToBeAlone     *bool       `json:"prefersToBeAlone"`
	SpecialNeeds         *string     `json:"specialNeeds"`
}

type BondedInfo struct {
	BondedWith []string `json:"bondedWith"`
	IsBonded   *bool    `json:"isBonded"`
}

type Medical struct {
	CurrentMedications   []string           `json:"currentMedications"`
	HealthConcerns       []string           `json:"healthConcerns"`
	Microchip            Microchip          `json:"microchip"`
	SpayedOrNeutered     *bool              `json:"spayedOrNeutered"`
	SpayedOrNeuteredDate *string            `json:"spayedOrNeuteredDate"`
	Surgeries            []MedicalProcedure `json:"surgeries"`
	Vaccinations         Vaccinations       `json:"vaccinations"`
	VaccinationsUpToDate *bool              `json:"vaccinationsUpToDate"`
}

type Microchip struct {
	MicrochipCompany *string `json:"microchipCompany"`
	MicrochipID      *string `json:"microchipID"`
	Microchipped     *bool   `json:"microchipped"`
}

type MedicalProcedure struct {
	Date  string  `json:"date"`
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Notes *string `json:"notes"`
}

type Vaccinations struct {
	Bordetella      *VaccineRecord  `json:"bordetella"`
	CanineDistemper *VaccineSeries  `json:"canineDistemper"`
	FelineDistemper *VaccineSeries  `json:"felineDistemper"`
	FelineLeukemia  *VaccineSeries  `json:"felineLeukemia"`
	Leptospira      *VaccineSeries  `json:"leptospira"`
	Other           []VaccineRecord `json:"other"`
	Rabies          *VaccineRecord  `json:"rabies"`
}

type VaccineRecord struct {
	DateAdministered string  `json:"dateAdministered"`
	ExpiresAt        *string `json:"expiresAt"`
	Veterinarian     *string `json:"veterinarian"`
}

type VaccineSeries struct {
	IsComplete bool           `json:"isComplete"`
	Round1     *VaccineRecord `json:"round1"`
	Round2     *VaccineRecord `json:"round2"`
	Round3     *VaccineRecord `json:"round3"`
}

type Descriptions struct {
	AdditionalInformation []string `json:"additionalInformation"`
	Behavioral            *string  `json:"behavioral"`
	Fun                   *string  `json:"fun"`
	Origin                *string  `json:"origin"`
	Primary               *string  `json:"primary"`
	SpecialNeeds          *string  `json:"specialNeeds"`
	Spotlight             *string  `json:"spotlight"`
}

type Details struct {
	EnvironmentType        *string `json:"environmentType"`
	IntakeDate             *string `json:"intakeDate"`
	PreferredPetLitterType *string `json:"preferredPetLitterType"`
	ShelterLocation        *string `json:"shelterLocation"`
	Status                 string  `json:"status"` // Mandatory in TS
}

type Adoption struct {
	AdoptedBy          *string      `json:"adoptedBy"`
	Date               *string      `json:"date"`
	NewAdoptedName     *string      `json:"newAdoptedName"`
	Photo              *Photo       `json:"photo"`
	Fee                *float64     `json:"fee"`
	SurveyCompleted    *bool        `json:"surveyCompleted"`
	AdopterContactInfo *ContactInfo `json:"adopterContactInfo"`
}

type Foster struct {
	EndDate           *string      `json:"endDate"`
	ParentName        *string      `json:"parentName"`
	ParentPhoto       *Photo       `json:"parentPhoto"`
	StartDate         *string      `json:"startDate"`
	FosterContactInfo *ContactInfo `json:"fosterContactInfo"`
}

type Returned struct {
	Date       *string `json:"date"`
	IsReturned bool    `json:"isReturned"`
	Reason     *string `json:"reason"`
}

type Sponsored struct {
	Amount      *float64 `json:"amount"`
	Date        *string  `json:"date"`
	IsSponsored bool     `json:"isSponsored"`
	SponsoredBy *string  `json:"sponsoredBy"`
}

type Photo struct {
	IsPrimary    bool    `json:"isPrimary"`
	ThumbnailURL *string `json:"thumbnailUrl"`
	UploadedAt   string  `json:"uploadedAt"`
	URL          string  `json:"url"`
}

type ProfileSettings struct {
	IsSpotlightFeatured       bool `json:"isSpotlightFeatured"`
	ShowAdditionalInformation bool `json:"showAdditionalInformation"`
	ShowMedicalHistory        bool `json:"showMedicalHistory"`
}

type ContactInfo struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}
