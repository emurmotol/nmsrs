package models

var (
	Sex            = []string{"Male", "Female"}
	CivilStatus    = []string{"Single", "Married", "Widowed", "Separated"}
	EmploymentType = []string{"Employed", "Unemployed"} // TODO: Duiplicate in struct EmploymentType must be EmploymentStatus
	Unemployed     = []string{"Actively looking for work", "Resigned", "Terminated/Laid off, local", "Terminated/Laid off, abroad"}
	Disability     = []string{"Visual impairment", "Hearing impairment", "Speech impairment", "Physically handicapped"}
)

type Applicant struct {
	PersonalInformation           PersonalInformation              `schema:"personal_information"`
	FormalEducation               []*FormalEducation               `schema:"formal_education"`
	ProfessionalLicence           []*ProfessionalLicence           `schema:"professional_licence"`
	Eligibility                   []*Eligibility                   `schema:"eligibility"`
	TrainingAndRelevantExperience []*TrainingAndRelevantExperience `schema:"training_and_relevant_experience"`
	CertificateOfCompetence       []*CertificateOfCompetence       `schema:"certificate_of_competence"`
	WorkExperience                WorkExperience                   `schema:"work_experience"`
	OtherSkills                   []string                         `schema:"other_skills"`
	CreatedAt                     string                           `schema:"created_at"`
	UpdatedAt                     string                           `schema:"updated_at"`
}

type PersonalInformation struct {
	FamilyName            string                `schema:"family_name"`
	GivenName             string                `schema:"given_name"`
	MiddleName            string                `schema:"middle_name"`
	PresentAddress        PresentAddress        `schema:"present_address"`
	Birthdate             string                `schema:"birthdate"`
	PlaceOfBirth          string                `schema:"place_of_birth"`
	Age                   int                   `schema:"age"`
	Sex                   string                `schema:"sex"`
	Height                float32               `schema:"height"`
	Weight                float32               `schema:"weight"`
	Religion              string                `schema:"religion"`
	CivilStatus           string                `schema:"civil_status"`
	LandlineNumber        string                `schema:"landline_number"`
	MobileNumber          string                `schema:"mobile_number"`
	EmailAddress          string                `schema:"email_address"`
	EmploymentStatus      EmploymentStatus      `schema:"employment_status"`
	PreferredOccupation   []string              `schema:"preferred_occupation"`
	PreferredWorkLocation PreferredWorkLocation `schema:"preferred_work_location"`
	Disability            []string              `schema:"disability"`
	Language              Language              `schema:"language"`
}

type PresentAddress struct {
	HouseNumber  string `schema:"house_number"`
	Street       string `schema:"street"`
	Subdivision  string `schema:"subdivision"`
	Barangay     string `schema:"barangay"`
	Municipality string `schema:"municipality"`
	City         string `schema:"city"`
}

type EmploymentStatus struct {
	IsEmployed bool       `schema:"is_employed"`
	Unemployed string     `schema:"unemployed"`
	Terminated Terminated `schema:"terminated"`
}

type Terminated struct {
	Loc     string `schema:"loc"`
	Country string `schema:"country"`
	Reason  string `schema:"reason"`
}

type PreferredWorkLocation struct {
	Local          string `schema:"local"`
	Overseas       string `schema:"overseas"`
	PassportNumber string `schema:"passport_number"`
	ExpiryDate     string `schema:"expiry_date"`
}

type Language struct {
	Native []string `schema:"native"`
	Other  []string `schema:"other"`
}

type FormalEducation struct {
	Grade         float32 `schema:"grade"`
	Course        string  `schema:"course"`
	School        string  `schema:"school"`
	YearGraduated int     `schema:"year_graduated"`
	LastAttended  int     `schema:"last_attended"`
}

type ProfessionalLicence struct {
	Title      string `schema:"title"`
	ExpiryDate string `schema:"expiry_date"`
}

type Eligibility struct {
	Title     string `schema:"title"`
	YearTaken string `schema:"year_taken"`
}

type TrainingAndRelevantExperience struct {
	Name                string `schema:"name"`
	SkillsAcquired      string `schema:"skills_acquired"`
	YearsOfExperience   int    `schema:"years_of_experience"`
	CertificateReceived string `schema:"certificate_received"`
	IssuedBy            string `schema:"issued_by"`
}

type CertificateOfCompetence struct {
	Certificate string  `schema:"certificate"`
	Rating      float32 `schema:"rating"`
	IssuedBy    string  `schema:"issued_by"`
	DateIssued  string  `schema:"date_issued"`
}

type WorkExperience struct {
	Local    []*Local    `schema:"local"`
	Overseas []*Overseas `schema:"overseas"`
}

type Local struct {
	CompanyName                string `schema:"company_name"`
	Address                    string `schema:"address"`
	From                       string `schema:"from"`
	To                         string `schema:"to"`
	PositionHead               string `schema:"position_head"`
	IsRelatedToFormalEducation bool   `schema:"is_related_to_formal_education"`
}

type Overseas struct {
	Position    string `schema:"position"`
	CompanyName string `schema:"company_name"`
	Type        string `schema:"type"`
	From        string `schema:"from"`
	To          string `schema:"to"`
}
