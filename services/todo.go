package services

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool      `json:"completed,omitempty" bson:"completed,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

var client *mongo.Client

func New(mongo *mongo.Client) Todo {
	client = mongo

	return Todo{}
}
