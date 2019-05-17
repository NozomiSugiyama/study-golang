package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
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

type Handler struct {
	userRepository *model.UserRepository
}

func NewUserHandler(repo *model.UserRepository) *Handler {
	h := new(Handler)
	h.userRepository = repo
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	print("handle /wep/users")
	switch r.Method {
	case http.MethodGet:
		h.GetUsers(w, r)
	case http.MethodPost:
		h.CreateUsers(w, r)
	// case http.MethodPut:
	// 	h.UpdateUser(w, r)
	// case http.MethodDelete:
	// 	h.DeleteUser(w, r)
	default:
		http.Error(w, "Only GET and POST methods are supported.", http.StatusBadRequest)
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepository.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

	res, err := json.Marshal(resultUsers)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}

func (h *Handler) CreateUsers(w http.ResponseWriter, r *http.Request) {

	var user UserCreate

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser := model.User{
		Name:        user.Name,
		Birthday:    user.Birthday,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	err = h.userRepository.CreateUser(&newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(User{
		newUser.ID,
		newUser.Name,
		newUser.Birthday,
		newUser.Email,
		newUser.PhoneNumber,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}
