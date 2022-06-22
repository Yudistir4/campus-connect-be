package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `gorm:"null" json:"name"`
	Profilepic    string `gorm:"null" json:"profilepic"`
	Email         string `gorm:"size:255;not null;unique" json:"email"`
	Password      string `gorm:"size:255;not null;" json:"password"`
	Bio           string `gorm:"null" json:"bio"`
	Link          string `gorm:"null" json:"link"`
	Whatsapp      string `gorm:"null" json:"whatsapp"`
	UserType      string `gorm:"null" json:"userType"`
	IdMahasiswa   int
	Mahasiswa     Mahasiswa `gorm:"foreignKey:IdMahasiswa"`
	IdOrganisasi  int
	Organisasi    Organisasi `gorm:"foreignKey:IdOrganisasi"`
	IdUniversitas int
	Universitas   Universitas `gorm:"foreignKey:IdUniversitas"`

	//Register
	// NameRektor string `gorm:"size:255;not null;" json:"name_rektor"`
	// KtpRektor  string `gorm:"size:255;not null;" json:"ktp_rektor"`
	// Isverified bool   `gorm:"default:false" json:"isverified"`
	// Alamat     string `gorm:"null" json:"alamat"`
}
