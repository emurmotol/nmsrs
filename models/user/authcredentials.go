package user

type AuthCredentials struct {
	Email    string `schema:"email" json:"email" bson:"email,omitempty" validate:"required,email"`
	Password string `schema:"password" json:"password" bson:"password,omitempty" validate:"required,min=6"`
}
