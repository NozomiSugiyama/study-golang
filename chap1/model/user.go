package model

import (
	"fmt"
	"math/rand"
	"time"
)

// User Resut User Struct
type User struct {
	ID          int    `json:"id`
	Name        string `json:"name"`
	Birthday    string `json:"birthday"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

// UserCreate Create User Struct
type UserCreate struct {
	ID          *int   `json:"id`
	Name        string `json:"name"`
	Birthday    string `json:"birthday"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

// UserUpdate Update User Struct
type UserUpdate struct {
	ID          int     `json:"id`
	Name        *string `json:"name"`
	Birthday    *string `json:"birthday"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
}

// CreateUser Create user
func CreateUser(session session, user UserCreate) (User, error) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int63())
	newUser := User{
		rand.Int(),
		user.Name,
		user.Birthday,
		user.Email,
		user.PhoneNumber,
	}
	session.store.users = append(session.store.users, newUser)
	return newUser, nil
}

// ListUsers List users from store
func ListUsers(session session, user UserCreate) ([]User, error) {
	users := session.store.users
	return users, nil
}
