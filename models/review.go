package models


// Declaring struct type for Review structure
type Review struct {
    Author string    `json:"author" bson:"author"`
    Rating uint8     `json:"rating" bson:"rating"`
    Date   string    `json:"date" bson:"date"`
    Body   string    `json:"body" bson:"body"`
}