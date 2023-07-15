package main

import (
	"github.com/AvinFajarF/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/user", controller.AddUser)
		v1.GET("/user", controller.ShowAllUsers)
		v1.PUT("/user/:id", controller.EditUser)
		v1.DELETE("/user/:id", controller.DeleteBook)
	}

	r.Run(":3000")

}