package api

import (
	"github.com/gin-gonic/gin"
	"my_mind/api/v1"
)

func Init() {
	r := gin.Default()
	
	r.GET("/user/:id", v1.GetUser)
	r.POST("/user", v1.CreateUser)
	r.GET("/users", v1.GetUsers)

	r.Run(":5002")
}