package data

import (
	"github.com/cconner57/adoption-os/backend/internal/validator"
)

type HouseholdMember struct {
	Age    string `json:"age"`
	Gender string `json:"gender"`
	Count  int    `json:"count"`
}

type SurrenderApplication struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	PhoneNumber   string `json:"phoneNumber"`
	Email         string `json:"email"`
	StreetAddress string `json:"streetAddress"`
	City          string `json:"city"`
	State         string `json:"state"`
	ZipCode       string `json:"zipCode"`

	WhenToSurrenderAnimal   string `json:"whenToSurrenderAnimal"`
	AnimalName              string `json:"animalName"`
	AnimalSex               string `json:"animalSex"`
	AnimalAge               string `json:"animalAge"`
	AnimalOwnershipDuration string `json:"animalOwnershipDuration"`
	AnimalLocationFound     string `json:"animalLocationFound"`
	AnimalWhySurrendered    string `json:"animalWhySurrendered"`

	HouseholdMembers     []HouseholdMember `json:"householdMembers"`
	OtherPetsInHousehold string            `json:"otherPetsInHousehold"`
	// Honeypot
	FaxNumber string `json:"fax_number"`

	// Behavior
	AnimalsBehaviorTowardsKnownPeople    string `json:"animalsBehaviorTowardsKnownPeople"`
	AnimalsBehaviorTowardsStrangers      string `json:"animalsBehaviorTowardsStrangers"`
	AnimalsBehaviorTowardsKnownAnimals   string `json:"animalsBehaviorTowardsKnownAnimals"`
	CommentsOnBehavior                   string `json:"commentsOnBehavior"`
	AnimalsReactionToNewPeople           string `json:"animalsReactionToNewPeople"`
	AnimalHouseTrained                   string `json:"animalHouseTrained"`
	AnimalSpendMajorityOfTime            string `json:"animalSpendMajorityOfTime"`
	AnimalLeftAloneDuration              string `json:"animalLeftAloneDuration"`
	AnimalWhenLeftAlone                  string `json:"animalWhenLeftAlone"`
	AnimalLeftAloneBehaviors             string `json:"animalLeftAloneBehaviors"`
	AnimalHowItPlays                     string `json:"animalHowItPlays"`
	AnimalToysItLikes                    string `json:"animalToysItLikes"`
	AnimalGamesItLikes                   string `json:"animalGamesItLikes"`
	AnimalScaredOfAnything               string `json:"animalScaredOfAnything"`
	AnimalScaredOfAnythingExplanation    string `json:"animalScaredOfAnythingExplanation"`
	AnimalBadHabits                      string `json:"animalBadHabits"`
	AnimalAllowedOnFurniture             string `json:"animalAllowedOnFurniture"`
	AnimalSleepAtNight                   string `json:"animalSleepAtNight"`
	AnimalBehaviorFoodOthers             string `json:"animalBehaviorFoodOthers"`
	AnimalBehaviorToysOthers             string `json:"animalBehaviorToysOthers"`
	AnimalProblemsRidingInCar            string `json:"animalProblemsRidingInCar"`
	AnimalProblemsRidingInCarExplanation string `json:"animalProblemsRidingInCarExplanation"`
	AnimalEscapedBefore                  string `json:"animalEscapedBefore"`
	AnimalEscapedBeforeExplanation       string `json:"animalEscapedBeforeExplanation"`

	// Aggression History
	AnimalEverAttackedPeople               string `json:"animalEverAttackedPeople"`
	AnimalEverAttackedPeopleExplanation    string `json:"animalEverAttackedPeopleExplanation"`
	AnimalEverAttackedOtherCats            string `json:"animalEverAttackedOtherCats"`
	AnimalEverAttackedOtherCatsExplanation string `json:"animalEverAttackedOtherCatsExplanation"`
	AnimalEverAttackedOtherDogs            string `json:"animalEverAttackedOtherDogs"`
	AnimalEverAttackedOtherDogsExplanation string `json:"animalEverAttackedOtherDogsExplanation"`

	// Medical
	AnimalVeterinarianList                       string `json:"animalVeterinarianList"`
	AnimalVeterinarianYearlyVisits               string `json:"animalVeterinarianYearlyVisits"`
	AnimalSpayedNeutered                         string `json:"animalSpayedNeutered"`
	AnimalVaccinationHistory                     string `json:"animalVaccinationHistory"`
	AnimalVaccinationsCurrent                    string `json:"animalVaccinationsCurrent"`
	AnimalTestedHeartworm                        string `json:"animalTestedHeartworm"`
	AnimalTestedHeartwormExplanation             string `json:"animalTestedHeartwormExplanation"`
	AnimalHeartwormPrevention                    string `json:"animalHeartwormPrevention"`
	AnimalHeartwormPreventionExplanation         string `json:"animalHeartwormPreventionExplanation"`
	AnimalMicrochipped                           string `json:"animalMicrochipped"`
	AnimalMicrochippedExplanation                string `json:"animalMicrochippedExplanation"`
	AnimalVetOrGroomerBehavior                   string `json:"animalVetOrGroomerBehavior"`
	AnimalVetMuzzled                             string `json:"animalVetMuzzled"`
	AnimalPastOrPresentHealthProblems            string `json:"animalPastOrPresentHealthProblems"`
	AnimalPastOrPresentHealthProblemsExplanation string `json:"animalPastOrPresentHealthProblemsExplanation"`
	AnimalCurrentMedications                     string `json:"animalCurrentMedications"`
	AnimalCurrentMedicationsExplanation          string `json:"animalCurrentMedicationsExplanation"`

	// Diet
	AnimalTypeOfFood            string `json:"animalTypeOfFood"`
	AnimalEatingFrequency       string `json:"animalEatingFrequency"`
	AnimalAmountOfFood          string `json:"animalAmountOfFood"`
	AnimalFoodTreats            string `json:"animalFoodTreats"`
	AnimalFoodTreatsExplanation string `json:"animalFoodTreatsExplanation"`

	AdditionalInformation string `json:"additionalInformation"`
}

func ValidateSurrenderApplication(v *validator.Validator, app *SurrenderApplication) {
	// Contact Info
	v.Check(app.FirstName != "", "firstName", "must be provided")
	v.Check(app.LastName != "", "lastName", "must be provided")
	v.Check(app.PhoneNumber != "", "phoneNumber", "must be provided")
	v.Check(app.Email != "", "email", "must be provided")
	v.Check(validator.Matches(app.Email, validator.EmailRX), "email", "must be a valid email address")
	v.Check(app.StreetAddress != "", "streetAddress", "must be provided")
	v.Check(app.City != "", "city", "must be provided")
	v.Check(app.State != "", "state", "must be provided")
	v.Check(app.ZipCode != "", "zipCode", "must be provided")

	// Animal Basic Info
	v.Check(app.AnimalName != "", "animalName", "must be provided")
	v.Check(app.AnimalAge != "", "animalAge", "must be provided")
	v.Check(app.AnimalSex != "", "animalSex", "must be provided")
	v.Check(app.AnimalOwnershipDuration != "", "animalOwnershipDuration", "must be provided")
	v.Check(app.AnimalLocationFound != "", "animalLocationFound", "must be provided")
	v.Check(app.AnimalWhySurrendered != "", "animalWhySurrendered", "must be provided")

	// Household
	v.Check(len(app.HouseholdMembers) > 0, "householdMembers", "must have at least one member")
	for _, member := range app.HouseholdMembers {
		v.Check(member.Count > 0, "householdMembers", "count must be positive")
	}
}
