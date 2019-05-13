package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name        string `json:"name"`
	Birthday    string `json:"birthday"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func main() {
	http.HandleFunc("/json/users", usersHandler)

	http.ListenAndServe(":8080", nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		response := []User{
			{"Nozomi Sugiyama", "1000/10/10", "Test@test.com", "09+0000-0000-0000"},
		}
		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(res))
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
