package controllers

import (
	"backend/database"
	"encoding/json"
	"net/http"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func GetTodoList(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT * FROM todos ORDER BY id DESC")
	if err != nil {
		http.Error(w, "db error", 500)
		return
	}
    defer rows.Close()

    var todos []Todo

	for rows.Next() {
		var t Todo
		rows.Scan(&t.ID, &t.Task, &t.Done)
		todos = append(todos, t)
	}

	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "invalid JSON", 400)
		return
	}

    result, err := database.DB.Exec(
		"INSERT INTO todos (task, done) VALUES (?, ?)",
		t.Task, t.Done,
	)
	if err != nil {
		http.Error(w, "db error", 500)
		return
	}

	id, _ := result.LastInsertId()
	t.ID = int(id)

	json.NewEncoder(w).Encode(t)
}

func UpdateTodoById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var update struct {
		Done bool `json:"done"`
	}

	json.NewDecoder(r.Body).Decode(&update)

	_, err := database.DB.Exec(
		"UPDATE todos SET done = ? WHERE id = ?",
		update.Done, id,
	)
	if err != nil {
		http.Error(w, "db error", 500)
		return
	}

	w.Write([]byte("updated"))
}
