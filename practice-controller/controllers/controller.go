package controllers

import (
    "fmt"
    "net/http"
)

func HelloController(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        id := r.PathValue("id")
        w.Write([]byte(fmt.Sprintf("Hello, user %s!", id)))
        return
    }
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte("Method not allowed"))
}