package registrant

type BasicInformation struct {
	StreetSubdivision string            `schema:"street_subdivision" mapstructure:"street_subdivision" json:"street_subdivision" bson:"StreetSubdivision,omitempty" validate:"required,min=1"`
	CityMunicipality  string            `schema:"city_municipality" mapstructure:"city_municipality" json:"city_municipality" bson:"CityMunicipality,omitempty" validate:"required"`
	Province          string            `schema:"province" mapstructure:"province" json:"province" bson:"province,omitempty" validate:"required"`
	Barangay          string            `schema:"barangay" mapstructure:"barangay" json:"barangay" bson:"barangay,omitempty" validate:"required"`
	PlaceOfBirth      string            `schema:"place_of_birth" mapstructure:"place_of_birth" json:"place_of_birth" bson:"placeOfBirth,omitempty" validate:"required,min=3"`
	Religion          string            `schema:"religion" mapstructure:"religion" json:"religion" bson:"religion,omitempty" validate:"required"`
	CivilStatus       map[string]string `schema:"civil_status" mapstructure:"civil_status" json:"civil_status" bson:"civilStatus,omitempty" validate:"required"`
	Sex               int               `schema:"sex" mapstructure:"sex" json:"sex" bson:"sex,omitempty" validate:"required"`
	Age               int               `schema:"age" mapstructure:"age" json:"age" bson:"age,omitempty" validate:"required,lte=122"`
	Height            float32           `schema:"height" mapstructure:"height" json:"height" bson:"height,omitempty" validate:"required"`
	Weight            float32           `schema:"weight" mapstructure:"weight" json:"weight" bson:"weight,omitempty" validate:"required"`
	Landline          string            `schema:"landline" mapstructure:"landline" json:"landline" bson:"landline,omitempty" validate:"required"`
	Mobile            string            `schema:"mobile" mapstructure:"mobile" json:"mobile" bson:"mobile,omitempty" validate:"required,min=12"`
	Email             string            `schema:"email" mapstructure:"email" json:"email" bson:"email,omitempty" validate:"required,email"`
}
