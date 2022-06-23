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
	idUserOrganisasi := c.Query("idUserOrganisasi")
	idUserUniversitas := c.Query("idUserUniversitas")
	name := c.Query("name")
	userType := c.Query("userType")
	isVerified, _ := strconv.ParseBool(c.Query("isVerified"))
	user := models.User{Name: name, UserType: userType}
	order := c.Query("order")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	startIndex := (page - 1) * limit

	var count int64 = 1

	var users []models.User

	if name != "" && userType == "" {
		fmt.Println("search")
		models.DB.
			Preload("Universitas").
			Preload("Mahasiswa").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Fakultas").
			Preload("Organisasi").
			Order(order).Limit(limit).Offset(startIndex).
			Where("name LIKE ?", "%"+name+"%").
			Find(&users)
	} else if idUserUniversitas != "" {
		fmt.Println("List Mahasiswa / Organisasi =====")

		IdUserUniversitas, _ := strconv.Atoi(idUserUniversitas)

		if userType == "mahasiswa" && name != "" {
			fmt.Println("mahasiswa name =====")

			models.DB.
				Preload("Mahasiswa").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Fakultas").
				Joins("left join mahasiswas on mahasiswas.id = users.id_mahasiswa").
				Order(order).Limit(limit).Offset(startIndex).
				Where("mahasiswas.id_user_universitas = ?", IdUserUniversitas).
				Where("user_type = ?", userType).Where("name LIKE ?", "%"+name+"%").
				Find(&users)
		} else if userType == "organisasi" && name != "" {
			fmt.Println("organisasi name =====")

			models.DB.
				Preload("Organisasi").
				Joins("left join organisasis on organisasis.id = users.id_organisasi").
				Order(order).Limit(limit).Offset(startIndex).
				Where("organisasis.id_user_universitas = ?", idUserUniversitas).
				Where("user_type = ?", userType).Where("name LIKE ?", "%"+name+"%").
				Find(&users)
		} else if userType == "mahasiswa" {
			models.DB.
				Preload("Mahasiswa").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Fakultas").
				Joins("left join mahasiswas on mahasiswas.id = users.id_mahasiswa").
				Order(order).Limit(limit).Offset(startIndex).
				Where("mahasiswas.id_user_universitas = ?", IdUserUniversitas).
				// Where("organisasis.id_user_universitas = ?", idUserUniversitas).
				Where("user_type = ?", userType).
				// Where("name LIKE ?", "%"+name+"%").
				Find(&users)

			models.DB.Table("users").Joins("left join mahasiswas on mahasiswas.id = users.id_mahasiswa").Where("mahasiswas.id_user_universitas = ?", IdUserUniversitas).Where("user_type = ?", userType).Count(&count)

		} else if userType == "organisasi" {
			models.DB.
				Preload("Organisasi").
				// Joins("left join mahasiswas on mahasiswas.id = users.id_mahasiswa").
				Joins("left join organisasis on organisasis.id = users.id_organisasi").
				Order(order).Limit(limit).Offset(startIndex).

				// Where("mahasiswas.id_user_universitas = ?", IdUserUniversitas).

				Where("organisasis.id_user_universitas = ?", idUserUniversitas).
				Where("user_type = ?", userType).
				// Where("name LIKE ?", "%"+name+"%").
				Find(&users)

			models.DB.Table("users").Joins("left join organisasis on organisasis.id = users.id_organisasi").Where("organisasis.id_user_universitas = ?", idUserUniversitas).Where("user_type = ?", userType).Count(&count)

		}

	} else if idUserOrganisasi != "" {
		fmt.Println("list mahasiswa in struktur organisasi")

		IdUserOrganisasi, _ := strconv.Atoi(idUserOrganisasi)

		if userType == "mahasiswa" && name != "" {

			models.DB.
				Preload("Mahasiswa").Preload("Mahasiswa.Jabatan").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Fakultas").
				Joins("left join mahasiswas on mahasiswas.id = users.id_mahasiswa").
				Order(order).Limit(limit).Offset(startIndex).
				Where("mahasiswas.id_user_universitas = ?", IdUserOrganisasi).
				Where("user_type = ?", userType).Where("name LIKE ?", "%"+name+"%").
				Find(&users)
		} else if userType == "mahasiswa" {
			models.DB.
				Preload("Mahasiswa").Preload("Mahasiswa.Jabatan").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Fakultas").
				Joins("left join mahasiswas on mahasiswas.id = users.id_mahasiswa").
				Joins("left join jabatans on jabatans.id = mahasiswas.id_jabatan").
				Order(order).Limit(limit).Offset(startIndex).
				Where("jabatans.id_user_organisasi = ?", IdUserOrganisasi).
				// Where("organisasis.id_user_universitas = ?", idUserUniversitas).
				Where("user_type = ?", userType).
				Find(&users)

			// models.DB.Table("users").Joins("left join mahasiswas on mahasiswas.id = users.id_mahasiswa").Where("mahasiswas.id_user_universitas = ?", IdUserOrganisasi).Where("user_type = ?", userType).Count(&count)

		}

	} else if c.Query("isVerified") != "" {
		// Get user berdasarkan table universitas dengan condisi isverified = true/false (blom jadi)
		fmt.Println("Verified")
		// models.DB.Raw("SELECT * FROM users t1 LEFT JOIN universitas t2 ON t1.id_universitas = t2.id WHERE t2.is_verified = ? LIMIT ? OFFSET ?", isVerified, limit, startIndex).Limit(1).Scan(&result)
		models.DB.
			Preload("Universitas").
			Joins("left join universitas on universitas.id = users.id_universitas").
			Where("user_type = ?", userType).
			Where("universitas.is_verified = ?", isVerified).
			Order(order).Limit(limit).Offset(startIndex).
			Find(&users)
	} else {
		fmt.Println("query")
		models.DB.Where(&user).Order(order).Limit(limit).Offset(startIndex).Preload("Mahasiswa").Preload("Organisasi").Preload("Universitas").Find(&users)
	}

	c.JSON(http.StatusOK, gin.H{

		"totalData":  len(users),
		"totalPages": math.Ceil(float64(count) / float64(limit)),
		"page":       page,
		"limit":      limit,
		"data":       users})
}
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	err := models.DB.Preload("Universitas").Preload("Mahasiswa").Preload("Mahasiswa.Jabatan").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Fakultas").
		Preload("Organisasi").First(&user, id).Error
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

type UserInput struct {
	ID            int
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
	IdFakultas      int
	IdProdi         int

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
		mahasiswa.IdFakultas = userInput.IdFakultas
		mahasiswa.IdProdi = userInput.IdProdi

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

// -------- JABATANS ------------

type InputJabatan struct {
	NamaJabatan      string
	NamaOrganisasi   string
	IdMahasiswa      int
	IdUserMahasiswa  int
	IdUserOrganisasi int
}

func CreateJabatan(c *gin.Context) {

	var user models.User
	var inputJabatan InputJabatan
	var mahasiswa models.Mahasiswa
	var jabatan models.Jabatan
	err := c.ShouldBindJSON(&inputJabatan)
	if err != nil {
		log.Fatal(err)
	}

	// Find Mahasiswa
	err = models.DB.First(&mahasiswa, inputJabatan.IdMahasiswa).Error
	if err != nil {
		log.Fatal("ERROR Find", err)
	}
	if mahasiswa.IdJabatan != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user sudah punya jabatan"})
		return
	}

	// create Jabatan
	jabatan.NamaJabatan = inputJabatan.NamaJabatan
	jabatan.NamaOrganisasi = inputJabatan.NamaOrganisasi
	jabatan.IdUserMahasiswa = inputJabatan.IdUserMahasiswa
	jabatan.IdUserOrganisasi = inputJabatan.IdUserOrganisasi
	err = models.DB.Create(&jabatan).Error
	if err != nil {
		log.Fatal("ERROR CREATE", err)
	}

	// assign ID jabatan to table Mahasiswa
	mahasiswa.IdJabatan = int(jabatan.ID)
	// save
	models.DB.Save(&mahasiswa)
	models.DB.Preload("Mahasiswa").Preload("Mahasiswa.Jabatan").First(&user, inputJabatan.IdUserMahasiswa)

	c.JSON(http.StatusOK, gin.H{"data": &user})

}

func UpdateJabatan(c *gin.Context) {}

// delete jabatan
func DeleteJabatan(c *gin.Context) {
	// Note

	id := c.Param("id")
	var user models.User
	var mahasiswa models.Mahasiswa
	var jabatan models.Jabatan

	// find user to get Id table mahasiswa
	err := models.DB.Preload("Mahasiswa").First(&user, id).Error
	if err != nil {
		log.Fatal("ERROR Find", err)
	}

	// find table mahasiswa
	err = models.DB.First(&mahasiswa, user.IdMahasiswa).Error
	if err != nil {
		log.Fatal("ERROR Find", err)
	}

	// Delete jabatan
	err = models.DB.Delete(&jabatan, mahasiswa.IdJabatan).Error
	if err != nil {
		log.Fatal("ERROR DELETE", err)
	}

	// update IdJabatan mahasiswa
	mahasiswa.IdJabatan = 0
	models.DB.Save(&mahasiswa)
	models.DB.Preload("Mahasiswa").Preload("Mahasiswa.Jabatan").First(&user, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete Jabatan",
		"data":    &user,
	})
}
