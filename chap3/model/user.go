package model

import (
	"github.com/jinzhu/gorm"
	"time"
)
type UserRepository struct {
	db *gorm.DB
}

type User struct {
	Model
	Name        string    `gorm:"size:255"`
	Birthday    time.Time `gorm:"size:255"`
	Email       string    `gorm:"size:255"`
	PhoneNumber string    `gorm:"size:255"`
}

const tableName = "users"

func NewUserRepository(db *gorm.DB) *UserRepository {
	repo := new(UserRepository)
	repo.db= db
	return repo
}

// CreateUser Create user
func (repo *UserRepository) CreateUser(user *User) error {
	repo.db.Table(tableName).Create(user)
	return nil
}

// ListUsers List users from store
func (repo *UserRepository) ListUsers() ([]User, error) {
	var users []User
	repo.db.Find(&users)
	return users, nil
}
