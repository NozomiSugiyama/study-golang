package controller

import (
	"time"

	"github.com/NozomiSugiyama/study-golang/chap3/model"
)

// User Resut User Struct
type User struct {
	ID          uint      `json:"id`
	Name        string    `json:"name"`
	Birthday    time.Time `json:"birthday"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
}

// UserCreate Create User Struct
type UserCreate struct {
	ID          *uint     `json:"id`
	Name        string    `json:"name"`
	Birthday    time.Time `json:"birthday"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
}

// UserUpdate Update User Struct
type UserUpdate struct {
	ID          uint       `json:"id`
	Name        *string    `json:"name"`
	Birthday    *time.Time `json:"birthday"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
}

func ListUsers(session *model.Session) ([]User, error) {
	users, err := model.ListUsers(session)
	if err != nil {
		return []User{}, err
	}

	resultUsers := make([]User, 0, len(users))

	for _, user := range users {
		resultUsers = append(
			resultUsers,
			User{
				user.ID,
				user.Name,
				user.Birthday,
				user.Email,
				user.PhoneNumber,
			},
		)
	}

	return resultUsers, nil
}

func CreateUser(session *model.Session, user UserCreate) (User, error) {
	newUser := model.User{
		Name:        user.Name,
		Birthday:    user.Birthday,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	err := model.CreateUser(session, &newUser)
	if err != nil {
		return User{}, err
	}

	resultUser := User{
		newUser.ID,
		newUser.Name,
		newUser.Birthday,
		newUser.Email,
		newUser.PhoneNumber,
	}
	return resultUser, nil
}
