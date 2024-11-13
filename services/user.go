package services

import (
	"gin-sample-project/dto"
	"gin-sample-project/models"
)

// FetchUserByID fetches a user by ID from the database
func FetchUserByID(id int) (*models.User, error) {
	return models.GetUserByID(id)
}

// RegisterUser creates a new user
func RegisterUser(payload *dto.User) error {

	var user models.User

	user.Email = payload.Email
	user.Username = payload.Username

	return models.CreateUser(&user)
}
