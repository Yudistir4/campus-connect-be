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

	router.POST("/users/:id/jabatan", controllers.CreateJabatan)
	router.PUT("/users/:id/jabatan", controllers.UpdateJabatan)
	router.DELETE("/users/:id/jabatan", controllers.DeleteJabatan)

	router.GET("/fakultas", controllers.GetAllFakultas)
	router.GET("/fakultas/:id", controllers.GetFakultas)
	router.POST("/fakultas", controllers.CreateFakultas)
	router.PUT("/fakultas", controllers.UpdateFakultas)
	router.DELETE("/fakultas/:id", controllers.DeleteFakultas)
	router.DELETE("/fakultas", controllers.DeleteAllFakultas)

	router.GET("/prodi", controllers.GetProdis)
	router.GET("/prodi/:id", controllers.GetProdi)
	router.POST("/prodi", controllers.CreateProdi)
	router.PUT("/prodi", controllers.UpdateProdi)
	router.DELETE("/prodi/:id", controllers.DeleteProdi)
	router.DELETE("/prodi", controllers.DeleteProdis)

	// router.GET("/posts", controllers.GetPosts)
	// router.GET("/posts/:id", controllers.GetPost)
	// router.POST("/posts", controllers.CreatePost)
	// router.DELETE("/posts/:id", controllers.DeletePost)
	// router.DELETE("/posts", controllers.DeletePosts)

	router.Run()
}
