package controllers

import (
	"first-app/models"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProdis(c *gin.Context) {
	idUserUniversitas := c.Query("idUserUniversitas")
	namaProdi := c.Query("namaProdi")
	order := c.Query("order")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	startIndex := (page - 1) * limit

	var count int64 = 1

	var prodis []models.Prodi

	if namaProdi != "" {
		models.DB.
			Preload("Fakultas").
			Order(order).Limit(limit).Offset(startIndex).
			Where("id_user_universitas = ?", idUserUniversitas).
			Where("nama_prodi LIKE ?", "%"+namaProdi+"%").
			Find(&prodis)

	} else {
		models.DB.
			Preload("Fakultas").
			Order(order).Limit(limit).Offset(startIndex).
			Where("id_user_universitas = ?", idUserUniversitas).
			Find(&prodis)
		models.DB.Table("prodis").
			Order(order).Limit(limit).Offset(startIndex).
			Where("id_user_universitas = ?", idUserUniversitas).
			Count(&count)

	}

	c.JSON(http.StatusOK, gin.H{

		"totalData":  len(prodis),
		"totalPages": math.Ceil(float64(count) / float64(limit)),
		"page":       page,
		"limit":      limit,
		"data":       prodis})
}

func GetProdi(c *gin.Context) {
	id := c.Param("id")
	var prodi models.Prodi
	err := models.DB.Preload("Fakultas").First(&prodi, id).Error
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": &prodi})
}

func CreateProdi(c *gin.Context) {

	var prodi models.Prodi
	err := c.ShouldBindJSON(&prodi)
	if err != nil {
		log.Fatal(err)
	}

	err = models.DB.Create(&prodi).Error
	if err != nil {
		log.Fatal("ERROR CREATE", err)
	}

	c.JSON(http.StatusOK, gin.H{"data": &prodi})
	// c.JSON(http.StatusOK, gin.H{"data": &prodi})

}

func UpdateProdi(c *gin.Context) {}

// delete prodi
func DeleteProdi(c *gin.Context) {
	// Note
	// if userType == universitas maka hapus seluruh prodi mahasiswa dan organisasi yang ada di univ tsb

	var prodi models.Prodi

	id := c.Param("id")
	err := models.DB.Delete(&prodi, id).Error
	if err != nil {
		log.Fatal("ERROR DELETE", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete ",
	})
}

// Delete All prodi (production ntr dihapus)
func DeleteProdis(c *gin.Context) {
	models.DB.Exec("DELETE FROM prodis")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
