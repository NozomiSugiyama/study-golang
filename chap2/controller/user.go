package controller

import (
	"github.com/NozomiSugiyama/study-golang/chap2/model"
)

func ListUsers() ([]model.User, error) {
	session := model.GetSession()
	users, err := model.ListUsers(session)
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}

func CreateUser(user model.UserCreate) (model.User, error) {
	session := model.GetSession()
	newUser, err := model.CreateUser(session, user)
	if err != nil {
		return model.User{}, err
	}
	return newUser, nil
}
