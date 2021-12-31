package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"my_mind/models"
	"my_mind/db"
)

func GetUsers(c *gin.Context) {
	users, err := db.GetUsers()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(c *gin.Context) {
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	user, err := db.GetUser(docID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	id, err := db.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}