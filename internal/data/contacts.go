package data

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CustomerModel struct {
	DB *mongo.Database
}

func (m CustomerModel) GetAll() []Customer {
	var customers []Customer

	collection := m.DB.Collection("customers")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return customers
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var customer Customer
		if err := cursor.Decode(&customer); err != nil {
			continue
		}
		customers = append(customers, customer)
	}

	return customers
}

func (m CustomerModel) GetByID(id string) *Customer {
	var customer Customer

	collection := m.DB.Collection("customers")
	err := collection.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&customer)
	if err != nil {
		return nil
	}

	return &customer
}

func (m CustomerModel) Insert(customer Customer) {
	collection := m.DB.Collection("customers")
	_, err := collection.InsertOne(context.TODO(), customer)
	if err != nil {
		// Handle error (in a real application, you might want to return an error)
		return
	}
}
