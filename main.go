package main

import (
	"log"
	"net/http"
	"html/template"
	"path/filepath"
	"os"
)

func main() {
	pattern := filepath.Join("/template/", "*.tpl")
	tpl := template.Must(template.ParseGlob(pattern))
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

var sex = []string{"Male", "Female"}
var civilStatus = []string{"Single", "Married", "Widowed", "Separated"}
var employmentStatus = []string{"Employed", "Unemployed"}
var unemployed = []string{"Actively looking for work", "Resigned", "Terminated/Laid off, local", "Terminated/Laid off, abroad"}
var disability = []string{"Visual impairment", "Hearing impairment", "Speech impairment", "Physically handicapped"}

type Applicant struct {
	PersonalInformation           PersonalInformation `json:"personal_information"`
	FormalEducation               []FormalEducation `json:"formal_education"`
	ProfessionalLicence           []ProfessionalLicence `json:"professional_licence"`
	Eligibility                   []Eligibility `json:"eligibility"`
	TrainingAndRelevantExperience []TrainingAndRelevantExperience `json:"training_and_relevant_experience"`
	CertificateOfCompetence       []CertificateOfCompetence `json:"certificate_of_competence"`
	WorkExperience                WorkExperience `json:"work_experience"`
	OtherSkills                   []string `json:"other_skills"`
	CreatedAt                     string `json:"created_at"`
	UpdatedAt                     string `json:"updated_at"`
}

type PersonalInformation struct {
	FamilyName            string `json:"family_name"`
	GivenName             string `json:"given_name"`
	MiddleName            string `json:"middle_name"`
	PresentAddress        PresentAddress `json:"present_address"`
	Birthdate             string `json:"birthdate"`
	PlaceOfBirth          string `json:"place_of_birth"`
	Age                   int `json:"age"`
	Sex                   string `json:"sex"`
	Height                float32 `json:"height"`
	Weight                float32 `json:"weight"`
	Religion              string `json:"religion"`
	CivilStatus           string `json:"civil_status"`
	LandlineNumber        string `json:"landline_number"`
	MobileNumber          string `json:"mobile_number"`
	EmailAddress          string `json:"email_address"`
	EmploymentStatus      EmploymentStatus `json:"employment_status"`
	PreferredOccupation   []string `json:"preferred_occupation"`
	PreferredWorkLocation PreferredWorkLocation `json:"preferred_work_location"`
	Disability            []string `json:"disability"`
	Language              Language `json:"language"`
}

type PresentAddress struct {
	HouseNumber  string `json:"house_number"`
	Street       string `json:"street"`
	Subdivision  string `json:"subdivision"`
	Barangay     string `json:"barangay"`
	Municipality string `json:"municipality"`
	City         string `json:"city"`
}

type EmploymentStatus struct {
	IsEmployed bool `json:"is_employed"`
	Unemployed string `json:"unemployed"`
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
	PassportNumber string `json:"passport_number"`
	ExpiryDate     string `json:"expiry_date"`
}

type Language struct {
	Native []string `json:"native"`
	Other  []string `json:"other"`
}

type FormalEducation struct {
	Grade         float32 `json:"grade"`
	Course        string `json:"course"`
	School        string `json:"school"`
	YearGraduated int `json:"year_graduated"`
	LastAttended  int `jaon:"last_attended"`
}

type ProfessionalLicence struct {
	Title      string `json:"title"`
	ExpiryDate string `json:"expiry_date"`
}

type Eligibility struct {
	Title     string `json:"title"`
	YearTaken string `json:"year_taken"`
}

type TrainingAndRelevantExperience struct {
	Name                string `json:"name"`
	SkillsAcquired      string `json:"skills_acquired"`
	YearsOfExperience   int `json:"years_of_experience"`
	CertificateReceived string `json:"certificate_received"`
	IssuedBy            string `json:"issued_by"`
}

type CertificateOfCompetence struct {
	Certificate string `json:"certificate"`
	Rating      float32 `json:"rating"`
	IssuedBy    string `json:"issued_by"`
	DateIssued  string `json:"date_issued"`
}

type WorkExperience struct {
	Local    []Local `json:"local"`
	Overseas []Overseas `json:"overseas"`
}

type Local struct {
	CompanyName                string `json:"company_name"`
	Address                    string `json:"address"`
	From                       string `json:"from"`
	To                         string `json:"to"`
	PositionHead               string `json:"position_head"`
	IsRelatedToFormalEducation bool `json:"is_related_to_formal_education"`
}

type Overseas struct {
	Position    string `json:"position"`
	CompanyName string `json:"company_name"`
	Type        string `json:"type"`
	From        string `json:"from"`
	To          string `json:"to"`
}

type User struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
