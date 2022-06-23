package controllers

import (
	"first-app/models"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllFakultas(c *gin.Context) {
	idUserUniversitas := c.Query("idUserUniversitas")
	namaFakultas := c.Query("namaFakultas")
	order := c.Query("order")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	startIndex := (page - 1) * limit

	var count int64 = 1

	var allFakultas []models.Fakultas

	if namaFakultas != "" {
		models.DB.
			Order(order).Limit(limit).Offset(startIndex).
			Where("id_user_universitas = ?", idUserUniversitas).
			Where("nama_fakultas LIKE ?", "%"+namaFakultas+"%").
			Find(&allFakultas)

	} else {
		models.DB.
			Order(order).Limit(limit).Offset(startIndex).
			Where("id_user_universitas = ?", idUserUniversitas).
			Find(&allFakultas)
		models.DB.Table("fakultas").
			Order(order).Limit(limit).Offset(startIndex).
			Where("id_user_universitas = ?", idUserUniversitas).
			Count(&count)

	}

	c.JSON(http.StatusOK, gin.H{

		"totalData":  len(allFakultas),
		"totalPages": math.Ceil(float64(count) / float64(limit)),
		"page":       page,
		"limit":      limit,
		"data":       allFakultas})
}

func GetFakultas(c *gin.Context) {
	id := c.Param("id")
	var fakultas models.Fakultas
	err := models.DB.First(&fakultas, id).Error
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": &fakultas})
}

func CreateFakultas(c *gin.Context) {

	var fakultas models.Fakultas
	err := c.ShouldBindJSON(&fakultas)
	if err != nil {
		log.Fatal(err)
	}

	err = models.DB.Create(&fakultas).Error
	if err != nil {
		log.Fatal("ERROR CREATE", err)
	}

	c.JSON(http.StatusOK, gin.H{"data": &fakultas})
	// c.JSON(http.StatusOK, gin.H{"data": &fakultas})

}

func UpdateFakultas(c *gin.Context) {}

// delete fakultas
func DeleteFakultas(c *gin.Context) {
	// Note
	// if userType == universitas maka hapus seluruh fakultas mahasiswa dan organisasi yang ada di univ tsb

	var fakultas models.Fakultas

	id := c.Param("id")
	err := models.DB.Delete(&fakultas, id).Error
	if err != nil {
		log.Fatal("ERROR DELETE", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete ",
	})
}

// Delete All fakultas (production ntr dihapus)
func DeleteAllFakultas(c *gin.Context) {
	models.DB.Exec("DELETE FROM fakultas")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
