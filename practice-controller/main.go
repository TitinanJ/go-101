package main

import (
	"fmt"
	"net/http"

	"httptesting/controllers"
)

func main() {
    http.HandleFunc("/users/{id}", controllers.HelloController)
	http.HandleFunc("/users", controllers.UserController)
    fmt.Println("Server is running on http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}