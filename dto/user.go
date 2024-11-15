package dto

// User represents the structure of a user
type User struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}
