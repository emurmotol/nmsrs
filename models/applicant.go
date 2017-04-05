package models

var (
	Sex            = []string{"Male", "Female"}
	CivilStatus    = []string{"Single", "Married", "Widowed", "Separated"}
	EmploymentType = []string{"Employed", "Unemployed"} // TODO: Duiplicate in struct EmploymentType must be EmploymentStatus
	Unemployed     = []string{"Actively looking for work", "Resigned", "Terminated/Laid off, local", "Terminated/Laid off, abroad"}
	Disability     = []string{"Visual impairment", "Hearing impairment", "Speech impairment", "Physically handicapped"}
)

type Applicant struct {
	PersonalInformation           PersonalInformation             `json:"personalInformation"`
	FormalEducation               []FormalEducation               `json:"formalEducation"`
	ProfessionalLicence           []ProfessionalLicence           `json:"professionalLicence"`
	Eligibility                   []Eligibility                   `json:"eligibility"`
	TrainingAndRelevantExperience []TrainingAndRelevantExperience `json:"trainingAndRelevantExperience"`
	CertificateOfCompetence       []CertificateOfCompetence       `json:"certificateOfCompetence"`
	WorkExperience                WorkExperience                  `json:"workExperience"`
	OtherSkills                   []string                        `json:"otherSkills"`
	CreatedAt                     string                          `json:"createdAt"`
	UpdatedAt                     string                          `json:"updatedAt"`
}

type PersonalInformation struct {
	FamilyName            string                `json:"familyName"`
	GivenName             string                `json:"givenName"`
	MiddleName            string                `json:"middleName"`
	PresentAddress        PresentAddress        `json:"presentAddress"`
	Birthdate             string                `json:"birthdate"`
	PlaceOfBirth          string                `json:"placeOfBirth"`
	Age                   int                   `json:"age"`
	Sex                   string                `json:"sex"`
	Height                float32               `json:"height"`
	Weight                float32               `json:"weight"`
	Religion              string                `json:"religion"`
	CivilStatus           string                `json:"civilStatus"`
	LandlineNumber        string                `json:"landlineNumber"`
	MobileNumber          string                `json:"mobileNumber"`
	EmailAddress          string                `json:"emailAddress"`
	EmploymentStatus      EmploymentStatus      `json:"employmentStatus"`
	PreferredOccupation   []string              `json:"preferredOccupation"`
	PreferredWorkLocation PreferredWorkLocation `json:"preferredWorkLocation"`
	Disability            []string              `json:"disability"`
	Language              Language              `json:"language"`
}

type PresentAddress struct {
	HouseNumber  string `json:"houseNumber"`
	Street       string `json:"street"`
	Subdivision  string `json:"subdivision"`
	Barangay     string `json:"barangay"`
	Municipality string `json:"municipality"`
	City         string `json:"city"`
}

type EmploymentStatus struct {
	IsEmployed bool       `json:"isEmployed"`
	Unemployed string     `json:"unemployed"`
	Terminated Terminated `json:"terminated"`
}

type Terminated struct {
	Loc     string `json:"loc"`
	Country string `json:"country"`
	Reason  string `json:"reason"`
}

type PreferredWorkLocation struct {
	Local          string `json:"local"`
	Overseas       string `json:"overseas"`
	PassportNumber string `json:"passportNumber"`
	ExpiryDate     string `json:"expiryDate"`
}

type Language struct {
	Native []string `json:"native"`
	Other  []string `json:"other"`
}

type FormalEducation struct {
	Grade         float32 `json:"grade"`
	Course        string  `json:"course"`
	School        string  `json:"school"`
	YearGraduated int     `json:"yearGraduated"`
	LastAttended  int     `jaon:"lastAttended"`
}

type ProfessionalLicence struct {
	Title      string `json:"title"`
	ExpiryDate string `json:"expiryDate"`
}

type Eligibility struct {
	Title     string `json:"title"`
	YearTaken string `json:"yearTaken"`
}

type TrainingAndRelevantExperience struct {
	Name                string `json:"name"`
	SkillsAcquired      string `json:"skillsAcquired"`
	YearsOfExperience   int    `json:"yearsOfExperience"`
	CertificateReceived string `json:"certificateReceived"`
	IssuedBy            string `json:"issuedBy"`
}

type CertificateOfCompetence struct {
	Certificate string  `json:"certificate"`
	Rating      float32 `json:"rating"`
	IssuedBy    string  `json:"issuedBy"`
	DateIssued  string  `json:"dateIssued"`
}

type WorkExperience struct {
	Local    []Local    `json:"local"`
	Overseas []Overseas `json:"overseas"`
}

type Local struct {
	CompanyName                string `json:"companyName"`
	Address                    string `json:"address"`
	From                       string `json:"from"`
	To                         string `json:"to"`
	PositionHead               string `json:"positionHead"`
	IsRelatedToFormalEducation bool   `json:"isRelatedToFormalEducation"`
}

type Overseas struct {
	Position    string `json:"position"`
	CompanyName string `json:"companyName"`
	Type        string `json:"type"`
	From        string `json:"from"`
	To          string `json:"to"`
}
