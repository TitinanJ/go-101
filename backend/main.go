package main

import (
	"fmt"
	"net/http"

	"backend/controllers"
	"backend/database"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	database.InitDB()
	
    http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetTodoList(w, r)
			return
		}
		if r.Method == http.MethodPost {
			controllers.CreateTodo(w, r)
			return
		}
	})

	http.HandleFunc("/todo/{id}", controllers.UpdateTodoById)

    fmt.Println("Server is running on http://localhost:3000")
    http.ListenAndServe(":3000", enableCORS(http.DefaultServeMux))
}
