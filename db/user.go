package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"my_mind/models"
)

var MODEL = "users"

//Create creating a task in a mongo or document db
func CreateUser(user *models.User) (primitive.ObjectID, error) {
	user.ID = primitive.NewObjectID()
	return Create[*models.User](user, MODEL)
}

func GetUsers() ([]models.User, error) {
	return GetList[models.User](MODEL)
}

func GetUser(id primitive.ObjectID) (models.User, error) {
	return Get[models.User](MODEL, id)
}