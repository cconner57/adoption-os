package data

// -------------------------------------------------------------------------
//  1. CORE IDENTIFIERS
// -------------------------------------------------------------------------

var Species = []string{
	"cat",
	"dog",
}

var Sexes = []string{
	"male",
	"female",
	"unknown",
}

var PetSizes = []string{
	"extra-small",
	"small",
	"medium",
	"large",
	"extra-large",
}

var AgeGroups = []string{
	"baby",
	"young",
	"adult",
	"senior",
}

// -------------------------------------------------------------------------
//  2. PHYSICAL & BEHAVIOR
// -------------------------------------------------------------------------

var CoatLengths = []string{
	"short",
	"medium",
	"long",
	"wire",
	"hairless",
}

var EnergyLevels = []string{
	"low",
	"medium",
	"high",
}

var Environments = []string{
	"outdoor",
	"indoor",
	"indoor/outdoor",
}

var Temperaments = []string{
	"affectionate",
	"anxious",
	"bossy",
	"curious",
	"hunter",
	"independent",
	"laid-back",
	"playful",
	"shy",
	"vocal",
}

// -------------------------------------------------------------------------
//  3. HEALTH & STATUS LOGS
// -------------------------------------------------------------------------

var EatingStatuses = []string{
	"normal",
	"low",
	"none",
	"excessive",
	"assisted-feeding",
}

var DrinkingStatuses = []string{
	"normal",
	"low",
	"none",
	"dehydrated",
}

var ActivityLogs = []string{
	"normal",
	"calm",
	"lethargic",
	"energetic",
	"hyperactive",
}

var UrineStatuses = []string{
	"normal",
	"none",
	"blood",
	"straining",
}

var DefecationStatuses = []string{
	"normal",
	"diarrhea",
	"none",
	"constipated",
}

var MedicalConcerns = []string{
	"allergies - flea",
	"allergies - food",
	"allergies - skin",
	"anemia",
	"asthma",
	"bladder infection",
	"cancer",
	"cystitis",
	"dental problems",
	"diabetes",
	"ear infections",
	"feline immunodeficiency (fiv)",
	"feline infectious peritonitis (fip)",
	"feline leukemia virus (felv)",
	"gastrointestinal issues",
	"heartworm disease",
	"hyperthyroidism",
	"kidney disease",
	"obesity",
	"upper respiratory infections",
}

var PetStatuses = []string{
	"adopted",
	"adoption-pending",
	"archived",
	"available",
	"foster",
	"hold",
	"intake",
}

// -------------------------------------------------------------------------
//  4. BREEDS
// -------------------------------------------------------------------------

var CatBreeds = []string{
	"Unknown",
	"Abyssinian", "Aegean", "American Bobtail", "American Curl", "American Polydactyl",
	"American Ringtail", "American Shorthair", "American Wirehair", "Aphrodite", "Arabian Mau",
	"Asian Hsorthair", "Asian Semi-Longhair", "Australian Mist", "Balinese", "Bambino",
	"Bengal", "Birman", "Bombay", "Bramble", "Brazilian Shorthair", "British Longhair",
	"British Semi-Longhair", "British Shorthair", "Burmese", "Burmilla", "California Spangled",
	"Chantilly", "Chartreux", "Chausie", "Cheetoh", "Colorpoint Shorthair", "Cornish Rex",
	"Coupari", "Cymric", "Desert Lynx", "Devon Rex", "Domestic Longhair", "Domestic Medium Hair",
	"Domestic Shorthair", "Donskoy", "Dragon Li", "Dwelf Cat", "Egyptian Mau", "European Burmese",
	"European Shorthair", "Exotic Shorthair", "Foldex Cat", "German Rex", "Havana", "Highlander",
	"Himalayan", "Japanese Bobtail", "Javanese", "Kashmir", "Khao Manee", "Kinkalow", "Korat",
	"Korean Shorthair", "Korn Ja", "Kurilian Bobtail", "Lambkin", "Laperm", "Lykoi",
	"Maine Coon Polydactyl", "Maine Coon", "Mandalay", "Manx", "Mekong Bobtail", "Minskin",
	"Minuet", "Munchkin", "Nebelung", "Norwegian Forest Cat", "Ocicat", "Ojos Azules",
	"Oriental Longhair", "Oriental Shorthair", "Owyhee Bob", "Peke-Faced", "Persian", "Peterbald",
	"Pixiebob", "Raas", "Ragamuffin", "Ragdoll", "Russian Blue", "Sam Sawet", "Savannah",
	"Scottish Fold", "Scottish Straight", "Selkirk Rex", "Serengeti", "Serrade Petit", "Siamese",
	"Siberian", "Singapura", "Skookum", "Snowshoe", "Sokoke", "Somali", "Sphynx", "Suphalak",
	"Tennessee Rex", "Thai", "Tonkinese", "Toybob", "Toyger", "Turkish Angora", "Turkish Van",
	"Ukrainian Levkoy", "York Chocolate",
}

var DogBreeds = []string{
	"Unknown",
	"Afghan Hound", "Airedale Terrier", "Akita", "Alaskan Malamute", "American Bulldog",
	"American Eskimo Dog", "American Foxhound", "American Pit Bull Terrier",
	"American Staffordshire Terrier", "Anatolian Shepherd Dog", "Australian Cattle Dog",
	"Australian Shepherd", "Australian Terrier", "Basenji", "Basset Hound", "Beagle",
	"Bernese Mountain Dog", "Bichon Frise", "Bloodhound", "Border Collie", "Border Terrier",
	"Boston Terrier", "Bouvier des Flandres", "Boxer", "Brittany", "Bull Terrier", "Bulldog",
	"Cairn Terrier", "Cane Corso", "Cardigan Welsh Corgi", "Cavalier King Charles Spaniel",
	"Chihuahua", "Chow Chow", "Cocker Spaniel", "Collie", "Dachshund", "Dalmatian",
	"Doberman Pinscher", "English Bulldog", "English Setter", "English Springer Spaniel",
	"French Bulldog", "German Shepherd", "German Shorthaired Pointer", "Giant Schnauzer",
	"Golden Retriever", "Great Dane", "Great Pyrenees", "Havanese", "Irish Setter",
	"Irish Wolfhound", "Italian Greyhound", "Jack Russell Terrier", "Japanese Chin", "Keeshond",
	"Labrador Retriever", "Lhasa Apso", "Maltese", "Miniature Pinscher", "Miniature Schnauzer",
	"Mixed Breed", "Newfoundland", "Norwegian Elkhound", "Old English Sheepdog", "Papillon",
	"Pekingese", "Pembroke Welsh Corgi", "Pit Bull", "Pointer", "Pomeranian", "Poodle",
	"Portuguese Water Dog", "Rhodesian Ridgeback", "Rottweiler", "Saint Bernard", "Samoyed",
	"Schipperke", "Scottish Terrier", "Shar Pei", "Shiba Inu", "Shih Tzu", "Siberian Husky",
	"Soft Coated Wheaten Terrier", "Staffordshire Bull Terrier", "Standard Poodle",
	"Tibetan Mastiff", "Vizsla", "Weimaraner", "West Highland White Terrier", "Whippet",
	"Yorkshire Terrier",
}

// -------------------------------------------------------------------------
//  5. HELPERS
// -------------------------------------------------------------------------

// IsPermittedValue checks if a value exists in a permitted list.
func IsPermittedValue(value string, permittedValues ...string) bool {
	for _, v := range permittedValues {
		if value == v {
			return true
		}
	}
	return false
}

// IsValidBreed checks if a breed string exists in either the Cat or Dog list
func IsValidBreed(breed string) bool {
	return IsPermittedValue(breed, CatBreeds...) || IsPermittedValue(breed, DogBreeds...)
}
