package models

import "database/sql"

// Pet represents the main database row
type Pet struct {
    ID              string          `json:"id"`
    Name            string          `json:"name"`
    Species         string          `json:"species"`
    Breed           string          `json:"breed"`
    Sex             string          `json:"sex"`
    Status          string          `json:"status"`
    Size            string          `json:"size"`
    Color           string          `json:"color"`
    CoatLength      string          `json:"coatLength"`
    DateOfBirth     sql.NullTime    `json:"dateOfBirth"`
    IntakeDate      sql.NullTime    `json:"intakeDate"`
    BehaviorData    BehaviorData    `json:"behaviorData"`
    MedicalData     MedicalData     `json:"medicalData"`
    DescriptionData DescriptionData `json:"descriptionData"`
    AdoptionData    AdoptionData    `json:"adoptionData"`
    FosterData      FosterData      `json:"fosterData"`
}

// JSONB Structs
type BehaviorData struct {
    EnergyLevel     string   `json:"energyLevel"`
    IsHouseTrained  bool     `json:"isHouseTrained"`
    IsGoodWithKids  bool     `json:"isGoodWithKids"`
    IsGoodWithCats  bool     `json:"isGoodWithCats"`
    IsGoodWithDogs  bool     `json:"isGoodWithDogs"`
    PrefersAlone    bool     `json:"prefersToBeAlone"`
    SpecialNeeds    string   `json:"specialNeeds"`
    PersonalityTags []string `json:"personalityTags"`
    BondedWith      []string `json:"bondedWith"`
    IsBonded        bool     `json:"isBonded"`
}

type MedicalData struct {
    VaccinationsUpToDate bool     `json:"vaccinationsUpToDate"`
    SpayedOrNeutered     bool     `json:"spayedOrNeutered"`
    Microchipped         bool     `json:"microchipped"`
    MicrochipID          string   `json:"microchipID"`
    MicrochipCompany     string   `json:"microchipCompany"`
    HealthConcerns       []string `json:"healthConcerns"`
    CurrentMedications   []string `json:"currentMedications"`
    HealthSummary        string   `json:"healthSummary"`
}

type DescriptionData struct {
    Primary        string   `json:"primary"`
    Spotlight      string   `json:"spotlight"`
    Behavioral     string   `json:"behavioral"`
    Origin         string   `json:"origin"`
    FunFact        string   `json:"fun"`
    AdditionalInfo []string `json:"additionalInformation"`
}

type AdoptionData struct {
    IsAdopted       bool    `json:"isAdopted"`
    Date            *string `json:"date"`
    Price           float64 `json:"price"`
    NewAdoptedName  string  `json:"newAdoptedName"`
    AdoptedBy       string  `json:"adoptedBy"`
    SurveyCompleted bool    `json:"surveyCompleted"`
}

type FosterData struct {
    BeingFostered bool    `json:"beingFostered"`
    StartDate     *string `json:"startDate"`
    EndDate       *string `json:"endDate"`
    ParentName    string  `json:"parentName"`
}
