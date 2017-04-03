package models

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
