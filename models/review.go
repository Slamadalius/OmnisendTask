package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Declaring struct type for Review structure
type Review struct {
    Id     primitive.ObjectID   `json:"_id" bson:"_id"`
    Author string               `json:"author" bson:"author"`
    Rating uint8                `json:"rating" bson:"rating"`
    Date   string               `json:"date" bson:"date"`
    Body   string               `json:"body" bson:"body"`
}