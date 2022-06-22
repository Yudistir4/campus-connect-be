package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Mahasiswa{})
	DB.AutoMigrate(&Organisasi{})
	DB.AutoMigrate(&Universitas{})
}
