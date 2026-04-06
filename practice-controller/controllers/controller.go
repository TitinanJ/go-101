package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func HelloController(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        id := r.PathValue("id")
        w.Write([]byte(fmt.Sprintf("Hello, user %s!", id)))
        return
    }
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte("Method not allowed"))
}

func UserController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("GET users"))
		return
	}

	if r.Method == http.MethodPost {
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		response := fmt.Sprintf("User %s created", user.Name)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(response))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed"))
}