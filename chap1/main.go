package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/NozomiSugiyama/study-golang/chap1/model"
)

func main() {
	http.HandleFunc("/json/users", usersHandler)

	http.ListenAndServe(":8080", nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	session := model.GetSession()
	switch r.Method {
	case "GET":
		usres, err := model.ListUsers(session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(usres)
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
