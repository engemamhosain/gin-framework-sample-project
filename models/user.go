package models

import "gin-sample-project/db"

type User struct {
	ID       int    `json:"id" bson:"_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// Add other fields as necessary
}

func GetUserByID(id int) (*User, error) {
	var user User
	err := db.DB.QueryRow("SELECT id, username, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser inserts a new user into the database
func CreateUser(user *User) error {
	_, err := db.DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", user.Username, user.Email)
	return err
}
