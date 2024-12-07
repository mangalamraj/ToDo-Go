package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo/db"
	"todo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	todo.ID = primitive.NewObjectID()
	todo.CreatedAt = time.Now()

	collection := db.GetCollection("todo_app", "todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, todo)
	if err != nil {
		http.Error(w, "Failed to insert to-do", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}

}
