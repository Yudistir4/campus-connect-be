package main

import (
	"first-app/controllers"
	"first-app/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	router := gin.Default()
	router.Static("/files", "./files")
	// Apply the middleware to the router (works with groups too)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.POST("/upload", controllers.UploadFile)
	// controllers.UploadFile()
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:id/profile", controllers.UpdateUserProfile)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	router.DELETE("/users", controllers.DeleteUsers)

	router.GET("/jabatan", controllers.GetJabatans)
	router.POST("/jabatan", controllers.CreateJabatan)
	router.PUT("/jabatan/:id", controllers.UpdateJabatan)
	router.DELETE("/jabatan/:id", controllers.DeleteJabatan)
	router.DELETE("/jabatan", controllers.DeleteJabatans)

	router.GET("/fakultas", controllers.GetAllFakultas)
	router.GET("/fakultas/:id", controllers.GetFakultas)
	router.POST("/fakultas", controllers.CreateFakultas)
	router.PUT("/fakultas/:id", controllers.UpdateFakultas)
	router.DELETE("/fakultas/:id", controllers.DeleteFakultas)
	router.DELETE("/fakultas", controllers.DeleteAllFakultas)

	router.GET("/prodi", controllers.GetProdis)
	router.GET("/prodi/:id", controllers.GetProdi)
	router.POST("/prodi", controllers.CreateProdi)
	router.PUT("/prodi/:id", controllers.UpdateProdi)
	router.DELETE("/prodi/:id", controllers.DeleteProdi)
	router.DELETE("/prodi", controllers.DeleteProdis)

	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPost)
	router.POST("/posts/:id/save", controllers.SavePost)
	router.POST("/posts/:id/unsave", controllers.UnsavePost)
	router.POST("/posts/:id/like", controllers.LikePost)
	router.POST("/posts/:id/unlike", controllers.UnlikePost)
	router.POST("/posts", controllers.CreatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
	router.DELETE("/posts", controllers.DeletePosts)

	router.GET("/comments", controllers.GetComments)
	router.POST("/comments", controllers.CreateComment)
	router.DELETE("/comments/:id", controllers.DeleteComment)

	router.GET("/save", controllers.GetSaves)
	router.GET("/like", controllers.GetLikes)
	router.Run()
}
