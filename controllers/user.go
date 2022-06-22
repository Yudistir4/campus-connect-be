package controllers

import (
	"first-app/models"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	name := c.Query("name")
	isVerified := c.Query("isVerified")
	user := models.User{Name: name}
	order := c.Query("order")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	startIndex := (page - 1) * limit

	var count int64
	models.DB.Table("users").Count(&count)
	totalPages := math.Ceil(float64(count) / float64(limit))

	var users []models.User
	err := models.DB.Find(&users).Error
	if err != nil {
		log.Fatal(err)
	}

	if name != "" {
		fmt.Println("search")
		models.DB.Where("name LIKE ?", "%"+name+"%").Order(order).Limit(limit).Offset(startIndex).Preload("Mahasiswa").Preload("Organisasi").Preload("Universitas").Find(&users)
	} else if isVerified != "" {
		// Get user berdasarkan table universitas dengan condisi isverified = true/false (blom jadi)
		fmt.Println("Verified")
		models.DB.Debug().Order(order).Limit(limit).Offset(startIndex).Preload("Universitas", isVerified+" = ?", true).Find(&users)
	} else {
		fmt.Println("query")
		models.DB.Where(&user).Order(order).Limit(limit).Offset(startIndex).Preload("Mahasiswa").Preload("Organisasi").Preload("Universitas").Find(&users)
	}

	// models.DB.Model(&users).Association("IdMahasiswa")
	c.JSON(http.StatusOK, gin.H{
		"totalDocs":  len(users),
		"totalPages": totalPages,
		"page":       page,
		"limit":      limit,
		"data":       users})
}
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	err := models.DB.First(&user, id).Error
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

type UserInput struct {
	Name          string `gorm:"null" json:"name"`
	Profilepic    string `gorm:"null" json:"profilepic"`
	Email         string `gorm:"size:255;not null;unique" json:"email"`
	Password      string `gorm:"size:255;not null;" json:"password"`
	Bio           string `gorm:"null" json:"bio"`
	Link          string `gorm:"null" json:"link"`
	Whatsapp      string `gorm:"null" json:"whatsapp"`
	UserType      string `gorm:"null" json:"userType"`
	IdMahasiswa   int
	IdOrganisasi  int
	IdUniversitas int
	// mahasiswa
	Semester        uint
	Nim             string
	StatusMahasiswa string

	// organisasi or mahasiswa
	IdUserUniversitas int
	Universitas       string

	// universitas
	NamaRektor string
	KtpRektor  string
	IsVerified bool
	Alamat     string
}

func CreateUser(c *gin.Context) {

	var universitas models.Universitas
	var mahasiswa models.Mahasiswa
	var organisasi models.Organisasi
	var user models.User
	var userInput UserInput
	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		log.Fatal(err)
	}

	if userInput.UserType == "mahasiswa" {

		mahasiswa.Semester = userInput.Semester
		mahasiswa.Nim = userInput.Nim
		mahasiswa.StatusMahasiswa = userInput.StatusMahasiswa
		mahasiswa.IdUserUniversitas = userInput.IdUserUniversitas
		mahasiswa.Universitas = userInput.Universitas

		err = models.DB.Create(&mahasiswa).Error
		if err != nil {
			log.Fatal("ERROR CREATE===", err)
		}

		user.Name = userInput.Name
		user.Email = userInput.Email
		//Generate random password
		user.Password = "1234"
		user.UserType = userInput.UserType
		user.IdMahasiswa = int(mahasiswa.ID)
		// Send password to email mahasiswa
		// TO DO ...
	} else if userInput.UserType == "organisasi" {

		organisasi.IdUserUniversitas = userInput.IdUserUniversitas
		organisasi.Universitas = userInput.Universitas

		err = models.DB.Create(&organisasi).Error
		if err != nil {
			log.Fatal("ERROR CREATE===", err)
		}

		user.Name = userInput.Name
		user.Email = userInput.Email
		//Generate random password
		user.Password = "1234"
		user.UserType = userInput.UserType
		user.IdOrganisasi = int(organisasi.ID)
		// Send password to email mahasiswa
		// TO DO ...
	} else if userInput.UserType == "universitas" {

		universitas.NamaRektor = userInput.NamaRektor
		universitas.KtpRektor = userInput.KtpRektor
		universitas.IsVerified = userInput.IsVerified
		universitas.Alamat = userInput.Alamat

		err = models.DB.Create(&universitas).Error
		if err != nil {
			log.Fatal("ERROR CREATE UNIVERSITAS", err)
		}

		user.Name = userInput.Name
		user.Email = userInput.Email
		user.Password = userInput.Password
		user.UserType = userInput.UserType
		user.IdUniversitas = int(universitas.ID)
	}

	err = models.DB.Create(&user).Error
	if err != nil {
		log.Fatal("ERROR CREATE===", err)
	}

	c.JSON(http.StatusOK, gin.H{"data": &user, "dataBefore": &userInput})
	// c.JSON(http.StatusOK, gin.H{"data": &user})

}
func UpdateUser(c *gin.Context) {}

// delete user
func DeleteUser(c *gin.Context) {
	// Note
	// if userType == universitas maka hapus seluruh user mahasiswa dan organisasi yang ada di univ tsb

	var user models.User

	id := c.Param("id")
	err := models.DB.Delete(&user, id).Error
	if err != nil {
		log.Fatal("ERROR DELETE", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete ",
	})
}

// Delete All user (production ntr dihapus)
func DeleteUsers(c *gin.Context) {
	models.DB.Exec("DELETE FROM users")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
