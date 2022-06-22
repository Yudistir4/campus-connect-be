package main

import (
	"first-app/controllers"
	"first-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	router := gin.Default()
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	router.DELETE("/users", controllers.DeleteUsers)

	// router.GET("/posts", controllers.GetPosts)
	// router.GET("/posts/:id", controllers.GetPost)
	// router.POST("/posts", controllers.CreatePost)
	// router.DELETE("/posts/:id", controllers.DeletePost)
	// router.DELETE("/posts", controllers.DeletePosts)

	router.Run()
}
