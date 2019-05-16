package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/NozomiSugiyama/study-golang/chap3/controller"
	"github.com/NozomiSugiyama/study-golang/chap3/model"
)

var session *model.Session

func main() {
	var err error
	session, err = model.GetSession(
		os.Getenv("APP_DB_ENDPOINT"),
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASSWORD"),
	)
	if err != nil {
		log.Print(err)
		return
	}
	defer session.Close()

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		return
	}
	http.HandleFunc("/wep/users", usersHandler)
	fcgi.Serve(l, nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	print("handle /wep/users")
	switch r.Method {
	case "GET":
		users, err := controller.ListUsers(session)
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
		var user controller.UserCreate

		ct := r.Header.Get("Content-Type")
		switch {
		case strings.HasPrefix(ct, "multipart/form-data"):
			user.Name = r.FormValue("name")
			user.Email = r.FormValue("email")
			// Datetime(String) to int type conversion
			t, err := time.Parse("2006-01-02", r.FormValue("birthday"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			user.Birthday = t
			user.PhoneNumber = r.FormValue("phone_number")
		case ct == "application/x-www-form-urlencoded":
			user.Name = r.FormValue("name")
			user.Email = r.FormValue("email")
			// Datetime(String) to timestamp type conversion
			t, err := time.Parse("2006-01-02", r.FormValue("birthday"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			user.Birthday = t
			user.PhoneNumber = r.FormValue("phone_number")
		case ct == "application/json":
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, "content-type is a required field", http.StatusBadRequest)
		}

		newUser, err := controller.CreateUser(session, user)
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

func toStringFromUser(user controller.User) string {
	return strconv.Itoa(int(user.ID)) + "," + user.Name + "," + user.Email + "," + user.Birthday.Format("2006-01-02") + "," + user.PhoneNumber
}
