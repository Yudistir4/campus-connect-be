package controllers

import (
	"first-app/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type users struct {
	ID         int    `json:"id"`
	Name       string `gorm:"null" json:"name"`
	ProfilePic string `gorm:"null" json:"profilePic"`
}

type responseComment struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Comment   string    `gorm:"not null" json:"comment"`
	IdPost    int       `gorm:"not null" json:"idPost"`
	IdUser    int       `gorm:"not null" json:"idUser"`
	User      users     `gorm:"foreignKey:IdUser" json:"user"`
}

func GetComments(c *gin.Context) {

	idPost := c.Query("idPost")
	order := c.Query("order")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	startIndex := (page - 1) * limit

	// var comments []models.Comment
	var comments []responseComment

	models.DB.Model(&models.Comment{}).
		Preload("User").
		Order(order).Limit(limit).Offset(startIndex).
		Where("id_post = ?", idPost).
		Find(&comments)

	c.JSON(http.StatusOK, gin.H{

		"totalData": len(comments),
		"page":      page,
		"limit":     limit,
		"data":      comments})
}

func CreateComment(c *gin.Context) {

	var post models.Post
	var comment models.Comment

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
		return

	}

	// cek postingan masih ada/tidak
	err = models.DB.Take(&post, comment.IdPost).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "Comment Failed, Post Not Found"})
		return
	}

	// create comment
	err = models.DB.Create(&comment).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "Comment Failed"})
		return
	}

	var respon responseComment
	err = models.DB.Model(&models.Comment{}).Preload("User").Take(&respon, comment.ID).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "Get Comment Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &respon})
	// c.JSON(http.StatusOK, gin.H{"data": &comment})

}

// func UpdateProdi(c *gin.Context) {}

// delete comment
func DeleteComment(c *gin.Context) {
	// Note
	// if userType == universitas maka hapus seluruh comment mahasiswa dan organisasi yang ada di univ tsb

	var comment models.Comment

	id := c.Param("id")
	err := models.DB.Delete(&comment, id).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "Delete Comment Failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete",
	})
}

// Delete All comment (production ntr dihapus)
func DeleteComments(c *gin.Context) {
	models.DB.Exec("DELETE FROM comments")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
