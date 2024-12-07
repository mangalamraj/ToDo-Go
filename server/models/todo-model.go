package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,"`
	Title     string             `json:"title" bson:"title"`
	Completed bool               `json:"completed" bson:"completed,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
