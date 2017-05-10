package registrant

type PersonalInformation struct {
	FamilyName string `schema:"family_name" mapstructure:"family_name" json:"family_name" bson:"familyName,omitempty" validate:"required,min=2"`
	GivenName  string `schema:"given_name" mapstructure:"given_name" json:"given_name" bson:"givenName,omitempty" validate:"required,min=2"`
	MiddleName string `schema:"middle_name" mapstructure:"middle_name" json:"middle_name" bson:"middleName,omitempty" validate:"required,min=2"`
	Birthdate  string `schema:"birthdate" mapstructure:"birthdate" json:"birthdate" bson:"birthdate,omitempty" validate:"required"`
	Password   string `schema:"password" mapstructure:"password" json:"password" bson:"password,omitempty" validate:"required,min=3"`
}
