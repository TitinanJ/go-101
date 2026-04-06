package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Address string `json:"address,omitempty"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// w.Write([]byte("Hello, World!"))

			alice := Person{Name: "Alice", Age: 30}
			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(alice)
			if err != nil {
				http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
				return
			}

			return
		}

		if r.Method == http.MethodPost {
			w.Write([]byte("Hey! You are posting something!"))
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))

	})

	http.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			id := r.PathValue("id")
			w.Write([]byte(fmt.Sprintf("Hello, user %s!", id)))
			return
		}
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from /hello"))
	})

	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}