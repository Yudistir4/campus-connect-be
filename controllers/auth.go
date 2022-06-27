package controllers

import (
	"first-app/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Signup struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	NamaRektor string `json:"namaRektor" binding:"required"`
	// UserType   string `json:"userType"`
	// IsVerified bool   `json:"isVerified"`
	KtpRektor string `json:"ktpRektor" binding:"required"`
	Alamat    string `json:"alamat" binding:"required"`
}

//signup user universitas
func SignupUser(c *gin.Context) {
	var register Signup
	var cekUser models.User
	err := c.BindJSON(&register)
	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Something went wrong",
		})
		return
	}

	// if email same error

	models.DB.Where("email = ?", register.Email).First(&cekUser)
	if cekUser.Email == register.Email {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Email already exists",
		})
		return
	}

	university := models.Universitas{
		NamaRektor: register.NamaRektor,
		KtpRektor:  register.KtpRektor,
		IsVerified: false,
		Alamat:     register.Alamat,
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Error hashing password",
			"error":       err,
		})
		return
	}

	models.DB.Create(&university)
	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Error Create Universitas",
			"error":       err,
		})
		return
	}

	user := models.User{
		Name:          register.Name,
		Email:         register.Email,
		Password:      string(hashedPassword),
		UserType:      "universitas",
		IdUniversitas: university.ID,
	}

	models.DB.Create(&user)

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "Successfully created user",
		"data":        user,
	})
}

// signin user
type Signin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SigninUser(c *gin.Context) {
	var signin Signin
	err := c.BindJSON(&signin)
	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Please fill all field",
			"error":       err})
		return
	}
	var user models.User

	// preload data
	err = models.DB.Preload("Organisasi").Preload("Mahasiswa").Preload("Universitas").Where("email = ?", signin.Email).First(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "User not found",
		})
		return
	}
	if user.UserType == "universitas" && user.Universitas.IsVerified == false {
		c.JSON(400, gin.H{
			"message": "User Not Verified",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signin.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Invalid Password",
			"error":       err})
		return
	}

	// // if password not match

	// create token for user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    user.Email,
		"password": user.Password,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Error signing token",
			"error":       err,
		})
		return
	}

	user.Password = ""

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "Successfully login",
		"data":        user,
		"token":       tokenString,
	})

}
