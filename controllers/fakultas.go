package controllers

import (
	"first-app/models"

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

	// var count int64 = 1

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

		// models.DB.Table("fakultas").
		// 	Order(order).Limit(limit).
		// 	Where("id_user_universitas = ?", idUserUniversitas).
		// 	Count(&count)

	}

	c.JSON(http.StatusOK, gin.H{

		"totalData": len(allFakultas),
		// "totalPages": math.Ceil(float64(count) / float64(limit)),
		// "count":      count,
		"page":  page,
		"limit": limit,
		"data":  allFakultas})
}

func GetFakultas(c *gin.Context) {
	id := c.Param("id")
	var fakultas models.Fakultas
	err := models.DB.First(&fakultas, id).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &fakultas})
}

func CreateFakultas(c *gin.Context) {

	var fakultas models.Fakultas
	err := c.ShouldBindJSON(&fakultas)
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
		return
	}

	err = models.DB.Create(&fakultas).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "Create Fakultas Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &fakultas})
	// c.JSON(http.StatusOK, gin.H{"data": &fakultas})

}

func UpdateFakultas(c *gin.Context) {
	id := c.Param("id")
	var fakultas models.Fakultas
	var inputFakultas models.Fakultas
	err := models.DB.First(&fakultas, id).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Fakultas Not Found ",
		})
	}
	err = c.ShouldBindJSON(&inputFakultas)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})
	}
	fakultas.NamaFakultas = inputFakultas.NamaFakultas
	err = models.DB.Save(&fakultas).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Update Fakultas Failed",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    &fakultas,
	})

}

// delete fakultas
func DeleteFakultas(c *gin.Context) {

	var fakultas models.Fakultas
	var prodi models.Prodi

	id := c.Param("id")
	err := models.DB.Delete(&fakultas, id).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
		return
	}
	err = models.DB.Where("id_fakultas = ?", id).Delete(&prodi).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
		return
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
