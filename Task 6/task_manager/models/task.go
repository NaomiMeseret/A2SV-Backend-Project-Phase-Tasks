package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title string ` bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Date string `bson:"date" json:"date"` //due-date eg:2025-03-04
	Status string `bson:"status" json:"status"`

}