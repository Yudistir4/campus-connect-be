package controllers

import (
	"first-app/models"
	 
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetJabatans(c *gin.Context) {
	var jabatan []models.Jabatan
	models.DB.Find(&jabatan)
	c.JSON(200, gin.H{"data": &jabatan})
}

func CreateJabatan(c *gin.Context) {

	var user models.User
	var inputJabatan models.Jabatan
	var mahasiswa models.Mahasiswa
	var jabatan models.Jabatan
	err := c.ShouldBindJSON(&inputJabatan)
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
		return
	}

	// create Jabatan
	jabatan.NamaJabatan = inputJabatan.NamaJabatan
	jabatan.NamaOrganisasi = inputJabatan.NamaOrganisasi
	jabatan.IdUserMahasiswa = inputJabatan.IdUserMahasiswa
	jabatan.IdMahasiswa = inputJabatan.IdMahasiswa
	jabatan.IdUserOrganisasi = inputJabatan.IdUserOrganisasi
	err = models.DB.Create(&jabatan).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "something went wrong"})
		return
	}

	// Find Mahasiswa
	err = models.DB.First(&mahasiswa, inputJabatan.IdMahasiswa).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
		return
	}
	if mahasiswa.IdJabatan != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user sudah punya jabatan"})
		return
	}
	// assign ID jabatan to table Mahasiswa
	mahasiswa.IdJabatan = int(jabatan.ID)
	// save
	models.DB.Save(&mahasiswa)
	models.DB.Preload("Mahasiswa").Preload("Mahasiswa.Fakultas").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Jabatan").First(&user, inputJabatan.IdUserMahasiswa)

	c.JSON(http.StatusOK, gin.H{"data": &user})

}

func UpdateJabatan(c *gin.Context) {
	// idUser := c.Param(":id")
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	var inputJabatan models.Jabatan
	var mahasiswa models.Mahasiswa
	var mahasiswaBaru models.Mahasiswa
	var jabatan models.Jabatan
	err := c.ShouldBindJSON(&inputJabatan)
	if err != nil {
		c.JSON(500, gin.H{"message": "something went wrong"})
		return
	}

	// Find Jabatan
	err = models.DB.First(&jabatan, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "jabatan not found"})
		return
	}

	if inputJabatan.IdUserMahasiswa != jabatan.IdUserMahasiswa {
		// hapus idJabatan di user mahasiswa lama

		err = models.DB.First(&mahasiswa, jabatan.IdMahasiswa).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "mahasiswa not found"})
			return
		}
		mahasiswa.IdJabatan = 0
		err = models.DB.Save(&mahasiswa).Error
		if err != nil {
			c.JSON(400, gin.H{"message": "something went wrong"})
			return
		}
		// tambah idJabatan di user baru (table mahasiswanya)
		err = models.DB.First(&mahasiswaBaru, inputJabatan.IdMahasiswa).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "mahasiswa baru not found"})
			return
		}
		mahasiswaBaru.IdJabatan = id
		err = models.DB.Save(&mahasiswaBaru).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "save mahasiswa fail"})
			return
		}

	}

	// update Jabatan
	jabatan.NamaJabatan = inputJabatan.NamaJabatan
	jabatan.IdUserMahasiswa = inputJabatan.IdUserMahasiswa
	jabatan.IdMahasiswa = inputJabatan.IdMahasiswa
	err = models.DB.Save(&jabatan).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}
	models.DB.Preload("Mahasiswa").Preload("Mahasiswa.Fakultas").Preload("Mahasiswa.Prodi").Preload("Mahasiswa.Jabatan").First(&user, inputJabatan.IdUserMahasiswa)

	c.JSON(http.StatusOK, gin.H{"data": &user})
}

// delete jabatan
func DeleteJabatan(c *gin.Context) {

	id := c.Param("id")
	var user models.User
	var mahasiswa models.Mahasiswa
	var jabatan models.Jabatan

	err := models.DB.Take(&jabatan, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Jabatan not found"})
		return
	}

	err = models.DB.Take(&mahasiswa, jabatan.IdMahasiswa).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}

	// Delete jabatan
	err = models.DB.Delete(&jabatan, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
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

func DeleteJabatans(c *gin.Context) {
	models.DB.Exec("DELETE FROM jabatans")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
