package controllers

import (
	"first-app/models"
	"first-app/utils"
	"fmt"

	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
			Where("user_type != ?", "admin").
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
				Order(order).Limit(limit).Offset(startIndex).
				Preload("Mahasiswa").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Fakultas").
				Joins("left join mahasiswas on mahasiswas.id = users.id_mahasiswa").
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
				Joins("left join jabatans on jabatans.id = mahasiswas.id_jabatan").
				Order(order).Limit(limit).Offset(startIndex).
				Where("jabatans.id_user_organisasi = ?", IdUserOrganisasi).
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
		models.DB.Where(&user).Order(order).Limit(limit).Offset(startIndex).Preload("Mahasiswa").Preload("Organisasi").Preload("Universitas").Where("user_type != ?", "admin").Find(&users)
	}

	c.JSON(http.StatusOK, gin.H{

		"totalData":  len(users),
		"totalPages": math.Ceil(float64(count) / float64(limit)),
		"page":       page,
		"limit":      limit,
		"data":       &users})
}
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	err := models.DB.Preload("Universitas").Preload("Mahasiswa").Preload("Mahasiswa.Jabatan").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Fakultas").
		Preload("Organisasi").First(&user, id).Error
	if err != nil {
		c.JSON(404, gin.H{"message": "user not found"})
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"data": user})
}

type UserInput struct {
	ID            int
	Name          string `gorm:"null" json:"name"`
	ProfilePic    string `gorm:"null" json:"profilePic"`
	Email         string `gorm:"size:255;not null;unique" json:"email"`
	Password      string `gorm:"size:255;not null;" json:"password"`
	Bio           string `gorm:"null" json:"bio"`
	Link          string `gorm:"null" json:"link"`
	Linkedin      string `gorm:"null" json:"linkedin"`
	Instagram     string `gorm:"null" json:"instagram"`
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
type PasswordInput struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func CreateUser(c *gin.Context) {

	var universitas models.Universitas
	var mahasiswa models.Mahasiswa
	var organisasi models.Organisasi
	var user models.User
	var userInput UserInput
	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
		return
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
			c.JSON(500, gin.H{"message": "Create mahasiswa failed"})
			return
		}

		user.IdMahasiswa = int(mahasiswa.ID)
		// Send password to email mahasiswa
		// TO DO ...
	} else if userInput.UserType == "organisasi" {

		organisasi.IdUserUniversitas = userInput.IdUserUniversitas
		organisasi.Universitas = userInput.Universitas

		err = models.DB.Create(&organisasi).Error
		if err != nil {
			c.JSON(500, gin.H{"message": "Create organisasi failed"})
			return
		}

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
			c.JSON(500, gin.H{"message": "Create Universitas failed"})
			return
		}

		user.IdUniversitas = int(universitas.ID)
	}
	user.Name = userInput.Name
	user.Email = userInput.Email
	user.Password = userInput.Password
	user.UserType = userInput.UserType

	if userInput.UserType == "mahasiswa" || userInput.UserType == "organisasi" {
		// generate password
		user.Password = "1234"
		// user.Password = utils.RandomString(6)
		// send password to email
		// utils.Sendmail(user.Email, user.Password)
	}
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Error hashing password",
			"error":       err,
		})
		return
	}
	user.Password = string(hashedPassword)

	err = models.DB.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{"message": "Create User Failed"})
	}
	err = models.DB.Preload("Universitas").Preload("Mahasiswa").Preload("Mahasiswa.Jabatan").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Fakultas").
		Preload("Organisasi").Take(&user, user.ID).Error
	if err != nil {
		c.JSON(404, gin.H{"message": "Update User Failed"})
		return
	}
	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"data": &user})
	// c.JSON(http.StatusOK, gin.H{"data": &user})

}
func UpdateUserProfile(c *gin.Context) {

	id := c.Param("id")
	var userInput UserInput
	var user models.User
	err := models.DB.Take(&user, id).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "User Not Found"})
		return
	}
	err = c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong", "error": err})
		return
	}

	user.ProfilePic = userInput.ProfilePic
	user.Bio = userInput.Bio
	user.Whatsapp = userInput.Whatsapp
	user.Link = userInput.Link
	user.Linkedin = userInput.Linkedin
	user.Instagram = userInput.Instagram
	err = models.DB.Save(&user).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Update User Failed"})
		return
	}

	c.JSON(200, gin.H{"message": "Update User succcess"})
}
func UpdatePassword(c *gin.Context) {

	id := c.Param("id")
	var passwordInput PasswordInput
	var user models.User

	err := models.DB.Take(&user, id).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "User Not Found"})
		return
	}

	err = c.ShouldBindJSON(&passwordInput)
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong", "error": err})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordInput.OldPassword))
	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Invalid old Password",
			"error":       err})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordInput.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Error hashing password",
			"error":       err,
		})
		return
	}
	user.Password = string(hashedPassword)
	err = models.DB.Save(&user).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Update Password Failed"})
		return
	}

	c.JSON(200, gin.H{"message": "Update Password succcess"})
}
func UpdateVerifiedUniversitas(c *gin.Context) {

	id := c.Param("id")
	var userInput UserInput
	var user models.User
	var universitas models.Universitas

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong", "error": err})
		return
	}

	err = models.DB.Take(&user, id).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "universitas not Found"})
		return
	}
	err = models.DB.Take(&universitas, user.IdUniversitas).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "universitas not Found"})
		return
	}
	universitas.IsVerified = userInput.IsVerified

	err = models.DB.Save(&universitas).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "Update universitas  Failed"})
		return
	}
	if userInput.IsVerified == true {
		utils.SendmailVerified(user.Email, user.Name)
	}

	c.JSON(200, gin.H{"message": "Update Password succcess"})
}
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var userInput UserInput
	var user models.User
	var mahasiswa models.Mahasiswa
	var organisasi models.Organisasi
	var universitas models.Universitas
	err := models.DB.Preload("Universitas").Preload("Organisasi").Preload("Mahasiswa").Take(&user, id).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "User Not Found"})
		return
	}
	err = c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong", "error": err})
		return
	}

	user.Name = userInput.Name
	user.Email = userInput.Email
	if user.Email != userInput.Email {
		// Generate random password
		// Send password to email mahasiswa
		// TO DO ...
		// If Fail send email return email not valid
		// hashing password
		user.Password = "1234"

	}
	user.UserType = userInput.UserType

	if userInput.UserType == "mahasiswa" {
		err = models.DB.Take(&mahasiswa, userInput.IdMahasiswa).Error
		if err != nil {
			c.JSON(500, gin.H{"error": "Mahasiswa not Found"})
			return
		}
		mahasiswa.Semester = userInput.Semester
		mahasiswa.Nim = userInput.Nim
		mahasiswa.StatusMahasiswa = userInput.StatusMahasiswa
		mahasiswa.IdUserUniversitas = userInput.IdUserUniversitas
		mahasiswa.Universitas = userInput.Universitas
		mahasiswa.IdFakultas = userInput.IdFakultas
		mahasiswa.IdProdi = userInput.IdProdi

		err = models.DB.Save(&mahasiswa).Error
		if err != nil {
			c.JSON(500, gin.H{"message": "Update Mahasiswa  Failed"})
			return
		}
		user.IdMahasiswa = mahasiswa.ID

	} else if userInput.UserType == "organisasi" {
		err = models.DB.Take(&organisasi, userInput.IdOrganisasi).Error
		if err != nil {
			c.JSON(404, gin.H{"error": "Organisasi not Found"})
			return
		}
		organisasi.IdUserUniversitas = userInput.IdUserUniversitas
		organisasi.Universitas = userInput.Universitas

		err = models.DB.Save(&organisasi).Error
		if err != nil {
			c.JSON(500, gin.H{"message": "Update Organisasi  Failed"})
			return
		}
		user.IdOrganisasi = organisasi.ID
	} else if userInput.UserType == "universitas" {
		err = models.DB.Take(&universitas, user.IdUniversitas).Error
		if err != nil {
			c.JSON(404, gin.H{"error": "universitas not Found"})
			return
		}
		universitas.IsVerified = userInput.IsVerified

		err = models.DB.Save(&universitas).Error
		if err != nil {
			c.JSON(500, gin.H{"message": "Update universitas  Failed"})
			return
		}

	}

	err = models.DB.Save(&user).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "Update User Failed"})
		return
	}
	// kembalikan data
	err = models.DB.Preload("Universitas").Preload("Mahasiswa").Preload("Mahasiswa.Jabatan").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Fakultas").
		Preload("Organisasi").Take(&user, user.ID).Error
	if err != nil {
		c.JSON(404, gin.H{"message": "Update User Failed"})
		return
	}
	user.Password = ""

	c.JSON(200, gin.H{"message": "Update User succcess", "data": &user})
}

// delete user
func DeleteUser(c *gin.Context) {
	// Note
	// if userType == universitas maka hapus seluruh user mahasiswa dan organisasi yang ada di univ tsb
	// delete recor mahasiswa/universitas/ organisasi
	var user models.User

	id := c.Param("id")
	err := models.DB.Delete(&user, id).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete",
	})
}

// Delete All user (production ntr dihapus)
func DeleteUsers(c *gin.Context) {
	models.DB.Exec("DELETE FROM users")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
