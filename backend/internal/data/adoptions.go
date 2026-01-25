package data

import (
	"time"

	"github.com/cconner57/adoption-os/backend/internal/validator"
)

type CurrentPet struct {
	Name             string `json:"name"`
	SpeciesBreedSize string `json:"speciesBreedSize"`
	Age              string `json:"age"`
	Source           string `json:"source"`
	SpayedNeutered   string `json:"spayedNeutered"`
	LikesDogs        string `json:"likesDogs"`
}

type PastPet struct {
	Name             string `json:"name"`
	SpeciesBreedSize string `json:"speciesBreedSize"`
	Age              string `json:"age"`
	Source           string `json:"source"`
	SpayedNeutered   string `json:"spayedNeutered"`
	PassedAwayReason string `json:"passedAwayReason"`
}

type Child struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type AdoptionApplication struct {
	// Metadata (Internal)
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Status    string    `json:"status"`

	// Honeypot
	FaxNumber string `json:"fax_number"`

	// General Section
	FirstName          string      `json:"firstName"`
	LastName           string      `json:"lastName"`
	Age                interface{} `json:"age"`
	PetID              *string     `json:"petId"`
	PetName            *string     `json:"petName"`
	SpouseFirstName    *string     `json:"spouseFirstName"`
	SpouseLastName     *string     `json:"spouseLastName"`
	RoommatesNames     []string    `json:"roommatesNames"`
	ChildrenNamesAges  []Child     `json:"childrenNamesAges"`
	Email              *string     `json:"email"`
	Address            *string     `json:"address"`
	AddressLine2       *string     `json:"addressLine2"`
	City               *string     `json:"city"`
	State              *string     `json:"state"`
	Zip                *string     `json:"zip"`
	PhoneNumber        *string     `json:"phoneNumber"`
	CellPhoneNumber    *string     `json:"cellPhoneNumber"`
	AdultMembersAgreed *string     `json:"adultMembersAgreed"`

	// Home Section
	HomeType                     *string     `json:"homeType"`      // e.g. House, Apt
	HomeOwnership                *string     `json:"homeOwnership"` // Own, Rent
	LandlordName                 *string     `json:"landlordName"`
	LandlordPhoneNumber          *string     `json:"landlordPhoneNumber"`
	AllowPets                    *string     `json:"allowPets"`
	BreedRestrictionsWeightLimit *string     `json:"breedRestrictionsWeightLimit"`
	MonthlyFee                   *string     `json:"monthlyFee"`
	Allergies                    *string     `json:"allergies"`
	PrimaryOwner                 interface{} `json:"primaryOwner"`
	YearsAtAddress               *string     `json:"yearsAtAddress"`
	PreviousAddress              *string     `json:"previousAddress"`
	ExpectToMove                 *string     `json:"expectToMove"`
	TravelPlan                   *string     `json:"travelPlan"`

	// New Cat Section
	CatAccess                *string `json:"catAccess"`
	CatIndoorOutdoor         *string `json:"catIndoorOutdoor"`
	CatPreferenceBreed       *string `json:"catPreferenceBreed"`
	CatPreferencePhysical    *string `json:"catPreferencePhysical"`
	CatPreferencePersonality *string `json:"catPreferencePersonality"`
	CatPreferenceNotWant     *string `json:"catPreferenceNotWant"`
	WhyInterested            *string `json:"whyInterested"`
	AdoptionReason           *string `json:"adoptionReason"`
	OwnCatBefore             *string `json:"ownCatBefore"`
	OwnKittenBefore          *string `json:"ownKittenBefore"`
	AlreadyHaveVeterinarian  *string `json:"alreadyHaveVeterinarian"`
	CatAllowedHomeArea       *string `json:"catAllowedHomeArea"`
	CatHomeAloneHours        *string `json:"catHomeAloneHours"`
	CatDisciplineType        *string `json:"catDisciplineType"`
	CatEscapeSteps           *string `json:"catEscapeSteps"`

	// Current Pets
	CurrentPets       []CurrentPet `json:"currentPets"`
	CurrentlyHavePets *string      `json:"currentlyHavePets"`

	// Past Pets
	PastPets      []PastPet `json:"pastPets"`
	OwnPetsBefore *string   `json:"ownPetsBefore"`

	// Other Section
	BredAnimalDescription   *string  `json:"bredAnimalDescription"`
	OwnedDeclawedOrDebarked *string  `json:"ownedDeclawedOrDebarked"`
	MovedWithPet            *string  `json:"movedWithPet"`
	OwnedSpecialNeedsPet    *string  `json:"ownedSpecialNeedsPet"`
	MobilityDevice          *string  `json:"mobilityDevice"`
	SurrenderConditions     []string `json:"surrenderConditions"`
	SurrenderPlan           *string  `json:"surrenderPlan"`
	FoodTypeBrand           *string  `json:"foodTypeBrand"`
	AffordVetCare           *string  `json:"affordVetCare"`
	AffordEmergencyCost     *string  `json:"affordEmergencyCost"`

	// Signatures
	AgreementSignature1 *string `json:"agreementSignature1"`
	AgreementSignature2 *string `json:"agreementSignature2"`
	AgreementSignature3 *string `json:"agreementSignature3"`
	SignatureData       *string `json:"signatureData"`
}

func ValidateAdoptionApplication(v *validator.Validator, app *AdoptionApplication) {
	// General Section
	v.Check(app.FirstName != "", "firstName", "must be provided")
	v.Check(app.LastName != "", "lastName", "must be provided")
	v.Check(app.Age != nil, "age", "must be provided")
	v.Check(app.Email != nil && *app.Email != "", "email", "must be provided")
	v.Check(app.Address != nil && *app.Address != "", "address", "must be provided")
	v.Check(app.City != nil && *app.City != "", "city", "must be provided")
	v.Check(app.State != nil && *app.State != "", "state", "must be provided")
	v.Check(app.Zip != nil && *app.Zip != "", "zip", "must be provided")
	v.Check(app.PhoneNumber != nil && *app.PhoneNumber != "", "phoneNumber", "must be provided")
	v.Check(app.AdultMembersAgreed != nil && *app.AdultMembersAgreed != "", "adultMembersAgreed", "must be provided")

	// Home Section
	if app.HomeType != nil {
		v.Check(*app.HomeType != "", "homeType", "must be provided")
	}
	if app.HomeOwnership != nil {
		v.Check(*app.HomeOwnership != "", "homeOwnership", "must be provided")

		// Conditional Validation for Renters
		if *app.HomeOwnership == "Rent" {
			v.Check(app.LandlordName != nil && *app.LandlordName != "", "landlordName", "must be provided for renters")
			v.Check(app.LandlordPhoneNumber != nil && *app.LandlordPhoneNumber != "", "landlordPhoneNumber", "must be provided for renters")
		}
	}

	// New Cat Section
	// (Most are technically optional in the Go struct pointers, but usually required by business logic.
	//  We'll enforce a few key ones or just let the frontend validation happen first.
	//  Below are some critical ones from previous conversation context)
	v.Check(app.AdoptionReason != nil && *app.AdoptionReason != "", "adoptionReason", "must be provided")

	// Past Pets
	v.Check(app.OwnPetsBefore != nil && *app.OwnPetsBefore != "", "ownPetsBefore", "must be provided")

	// Other Section - Surrender related
	v.Check(app.SurrenderPlan != nil && *app.SurrenderPlan != "", "surrenderPlan", "must be provided")
	// v.Check(len(app.SurrenderConditions) > 0, "surrenderConditions", "must be provided") // Checkbox array, might be empty if none apply? Assuming at least one must be checked based on UI.

	// Signatures
	v.Check(app.AgreementSignature1 != nil && *app.AgreementSignature1 != "", "agreementSignature1", "must be provided")
	v.Check(app.AgreementSignature2 != nil && *app.AgreementSignature2 != "", "agreementSignature2", "must be provided")
	// Signature 3 Disabled based on user request
	// v.Check(app.AgreementSignature3 != nil && *app.AgreementSignature3 != "", "agreementSignature3", "must be provided")
	v.Check(app.SignatureData != nil && *app.SignatureData != "", "signatureData", "must be provided")
}
