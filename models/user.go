package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       	primitive.ObjectID  	`json:"id" bson:"_id,omitempty"`
	Username 	string 	`json:"username"`
	Firstname string 	`json:"firstname"`
	Lastname 	string 	`json:"lastname"`
	Email 		string 	`json:"email"`
	Age 			int 		`json:"age"`
}