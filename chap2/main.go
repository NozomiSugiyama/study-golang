package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"strconv"
	"strings"

	"github.com/NozomiSugiyama/study-golang/chap2/controller"
	"github.com/NozomiSugiyama/study-golang/chap2/model"
)

func main() {
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		return
	}
	http.HandleFunc("/wep/users", usersHandler)
	fcgi.Serve(l, nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := controller.ListUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		switch r.URL.Query().Get("format") {
		case "html":
			w.Header().Set("Content-Type", "text/html")
			for _, value := range users {
				fmt.Fprint(w, toStringFromUser(value)+"\n")
			}
		case "plain":
			w.Header().Set("Content-Type", "text/plain")
			for _, value := range users {
				fmt.Fprint(w, toStringFromUser(value)+"\n")
			}
		// json
		default:
			res, err := json.Marshal(users)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(res))
		}
	case "POST":
		// var user model.UserCreate{}
		user := &model.UserCreate{}

		ct := r.Header.Get("Content-Type")
		switch {
		case strings.HasPrefix(ct, "multipart/form-data"):
			user.Name = r.FormValue("name")
			user.Email = r.FormValue("email")
			user.Birthday = r.FormValue("birthday")
			user.PhoneNumber = r.FormValue("phone_number")
		case ct == "application/x-www-form-urlencoded":
			user.Name = r.FormValue("name")
			user.Email = r.FormValue("email")
			user.Birthday = r.FormValue("birthday")
			user.PhoneNumber = r.FormValue("phone_number")
		case ct == "application/json":
			decoder := json.NewDecoder(r.Body)
			// err := decoder.Decode(&user)
			err := decoder.Decode(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, "content-type is a required field", http.StatusBadRequest)
		}

		newUser, err := controller.CreateUser(*user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		switch r.URL.Query().Get("format") {
		case "html":
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, toStringFromUser(newUser)+"\n")
		case "plain":
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprint(w, toStringFromUser(newUser)+"\n")
		// json
		default:
			res, err := json.Marshal(newUser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(res))
		}
	default:
		fmt.Fprintf(w, "Only GET and POST methods are supported.")
	}
}

func toStringFromUser(user model.User) string {
	return strconv.Itoa(user.ID) + "," + user.Name + "," + user.Email + "," + user.Birthday + "," + user.PhoneNumber
}
