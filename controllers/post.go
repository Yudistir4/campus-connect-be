package controllers

import (
	"first-app/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Posts struct {
	ID                int       `json:"id"`
	CreatedAt         time.Time `json:"createdAt"`
	IdUser            int       `json:"idUser"`
	IdUserUniversitas int       `json:"idUserUniversitas"`
	User              users     `gorm:"foreignKey:IdUser" json:"user"`
	Materi            string    `json:"materi"`
	Caption           string    `json:"caption"`
	IsNews            bool      `json:"isNews"`
	IsSaved           bool      `json:"isSaved"`
	IsLiked           bool      `json:"isLiked"`
	JumlahLike        int       `json:"jumlahLike"`
	JumlahComment     int       `json:"jumlahComment"`
}

func GetPosts(c *gin.Context) {
	idUserUniversitas := c.Query("idUserUniversitas")
	idUser, _ := strconv.Atoi(c.Query("idUser"))
	isNews := c.Query("isNews")
	isSave := c.Query("isSave")

	order := c.Query("order")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	startIndex := (page - 1) * limit

	var posts []Posts
	var saves []models.Save

	if idUserUniversitas != "" {
		models.DB.
			Preload("User").
			Order(order).Limit(limit).Offset(startIndex).
			Where("id_user_universitas = ?", idUserUniversitas).
			Find(&posts)

	} else if isNews != "" {
		isNews, _ := strconv.ParseBool(isNews)
		models.DB.
			Preload("User").
			Order(order).Limit(limit).Offset(startIndex).
			Where("is_news = ?", isNews).
			Find(&posts)

	} else if isSave != "" {

		models.DB.
			Preload("Post").Preload("Post.User").Where("id_user = ?", idUser).
			Order(order).Limit(limit).Offset(startIndex).
			Find(&saves)
		for _, val := range saves {
			posts = append(posts, Posts{
				ID:                int(val.Post.ID),
				CreatedAt:         val.Post.CreatedAt,
				Materi:            val.Post.Materi,
				Caption:           val.Post.Caption,
				IdUserUniversitas: val.Post.IdUserUniversitas,
				IsNews:            val.Post.IsNews,
				IdUser:            val.Post.IdUser,
				User: users{ID: int(val.Post.User.ID),
					Name:       val.Post.User.Name,
					ProfilePic: val.Post.User.ProfilePic,
				},
			})
		}

	} else {
		models.DB.Preload("User").Order(order).Limit(limit).Offset(startIndex).Find(&posts)
	}

	for i, val := range posts {
		var save models.Save
		var like models.Like
		var countComment int64
		var countLike int64
		models.DB.Where("id_user = ? AND id_post = ?", idUser, val.ID).Take(&save)
		models.DB.Where("id_user = ? AND id_post = ?", idUser, val.ID).Take(&like)
		models.DB.Model(&models.Comment{}).Where("id_post = ?", val.ID).Count(&countComment)
		models.DB.Model(&models.Like{}).Where("id_post = ?", val.ID).Count(&countLike)

		if save.ID != 0 {
			posts[i].IsSaved = true
		}
		if like.ID != 0 {
			posts[i].IsLiked = true
		}
		posts[i].JumlahComment = int(countComment)
		posts[i].JumlahLike = int(countLike)

	}

	c.JSON(http.StatusOK, gin.H{

		"totalData": len(posts),

		"page":  page,
		"limit": limit,
		"data":  posts})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	err := models.DB.Preload("User").First(&post, id).Error
	if err != nil {
		c.JSON(404, gin.H{"message": "Post Not Found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": &post})
}
func SavePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	// cek post masih ada atau tidak
	err := models.DB.First(&post, id).Error
	if err != nil {
		c.JSON(404, gin.H{"message": "Post has been deleted"})
		return
	}

	var save models.Save

	err = c.ShouldBindJSON(&save)
	if err != nil {
		c.JSON(500, gin.H{"error": "something went wrong"})
		return
	}

	// cek udh di save atau blom
	err = models.DB.Where("id_post = ? AND id_user = ?", save.IdPost, save.IdUser).First(&save).Error
	if save.ID != 0 {
		c.JSON(400, gin.H{"message": "Post has been saved earlier "})
		return
	}

	// save
	err = models.DB.Create(&save).Error
	if err != nil {
		c.JSON(404, gin.H{"message": "failed save post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success save post"})
}

func UnsavePost(c *gin.Context) {

	var save models.Save
	err := c.ShouldBindJSON(&save)
	if err != nil {
		c.JSON(500, gin.H{"error": "something went wrong"})
		return
	}
	// unsave
	err = models.DB.Where("id_post = ? AND id_user = ?", save.IdPost, save.IdUser).Delete(&save).Error
	if err != nil {
		c.JSON(404, gin.H{"message": "failed unsave post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success unsave post"})
}
func LikePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	// cek post masih ada atau tidak
	err := models.DB.First(&post, id).Error
	if err != nil {
		c.JSON(404, gin.H{"message": "Post has been deleted"})
		return
	}

	var like models.Like

	err = c.ShouldBindJSON(&like)
	if err != nil {
		c.JSON(500, gin.H{"error": "something went wrong"})
		return
	}

	// cek udh di like atau blom
	err = models.DB.Where("id_post = ? AND id_user = ?", like.IdPost, like.IdUser).First(&like).Error
	if like.ID != 0 {
		c.JSON(400, gin.H{"message": "Post has been liked earlier "})
		return
	}

	// like
	err = models.DB.Create(&like).Error
	if err != nil {
		c.JSON(400, gin.H{"message": "failed like post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success like post"})
}

func UnlikePost(c *gin.Context) {

	var like models.Like
	err := c.ShouldBindJSON(&like)
	if err != nil {
		c.JSON(500, gin.H{"error": "something went wrong"})
		return
	}
	// unlike
	err = models.DB.Where("id_post = ? AND id_user = ?", like.IdPost, like.IdUser).Delete(&like).Error
	if err != nil {
		c.JSON(404, gin.H{"message": "failed unlike post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success unlike post"})
}

func GetSaves(c *gin.Context) {

	var save []models.Save

	// unsave
	err := models.DB.Find(&save).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "failed get saved"})
	}
	c.JSON(http.StatusOK, gin.H{"data": &save})
}
func GetLikes(c *gin.Context) {

	var like []models.Like

	// unlike
	err := models.DB.Find(&like).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "failed get liked"})
	}
	c.JSON(http.StatusOK, gin.H{"data": &like})
}

//create post
func CreatePost(c *gin.Context) {

	var post models.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = models.DB.Create(&post).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": post})
}

// delete post
func DeletePost(c *gin.Context) {
	// Delete all save yg terkait dengan postingan ini
	// TODO...

	var post models.Post

	id := c.Param("id")
	err := models.DB.Delete(&post, id).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete post",
	})
}

// Delete All post (production ntr dihapus)
func DeletePosts(c *gin.Context) {
	models.DB.Exec("DELETE FROM posts")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
