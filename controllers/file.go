package controllers

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {

	file, _ := c.FormFile("file")
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext

	path := "files/" + filename
	err := c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
	}
	c.JSON(200, gin.H{"filename": &filename})
}
