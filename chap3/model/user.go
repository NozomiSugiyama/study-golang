package model

import (
	"time"
)

type User struct {
	Model
	Name        string    `gorm:"size:255"`
	Birthday    time.Time `gorm:"size:255"`
	Email       string    `gorm:"size:255"`
	PhoneNumber string    `gorm:"size:255"`
}

const tableName = "users"

// CreateUser Create user
func CreateUser(session *Session, user *User) error {
	session.Table(tableName).Create(user)
	return nil
}

// ListUsers List users from store
func ListUsers(session *Session) ([]User, error) {
	var users []User
	session.Find(&users)
	return users, nil
}
