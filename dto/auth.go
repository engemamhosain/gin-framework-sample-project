package dto

// User represents the structure of a user
type Auth struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required"`
}
