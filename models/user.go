package models

import (
	"fmt"
	"gin-sample-project/db"
)

type User struct {
	ID       int    `json:"id" bson:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	// Add other fields as necessary
}

func GetUserByID(id int) (*User, error) {
	var user User
	err := db.DB.QueryRow("SELECT id, username, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUserName(username string) (*User, error) {

	var user User
	err := db.DB.QueryRow("SELECT id, username, email FROM users WHERE username = ? limit 1", username).Scan(&user.ID, &user.Username, &user.Email)
	fmt.Println("user ", user)
	fmt.Println("err ", err)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser inserts a new user into the database
func CreateUser(user *User) error {
	_, err := db.DB.Exec("INSERT INTO users (username,password,email) VALUES (?,?,?)", user.Username, user.Password, user.Email)
	return err
}
