package models

import "time"

type User struct {
	ID            int         `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time   `json:"createdAt"`
	Name          string      `gorm:"null" json:"name"`
	ProfilePic    string      `gorm:"null" json:"profilePic"`
	Email         string      `gorm:"size:255;not null;unique" json:"email"`
	Password      string      `gorm:"size:255;not null;" json:"password"`
	Bio           string      `gorm:"null" json:"bio"`
	Link          string      `gorm:"null" json:"link"`
	Instagram     string      `gorm:"null" json:"instagram"`
	Linkedin      string      `gorm:"null" json:"linkedin"`
	Whatsapp      string      `gorm:"null" json:"whatsapp"`
	UserType      string      `gorm:"null" json:"userType"`
	IdMahasiswa   int         `gorm:"null" json:"idMahasiswa"`
	Mahasiswa     Mahasiswa   `gorm:"foreignKey:IdMahasiswa" json:"mahasiswa"`
	IdOrganisasi  int         `gorm:"null" json:"idOrganisasi"`
	Organisasi    Organisasi  `gorm:"foreignKey:IdOrganisasi" json:"organisasi"`
	IdUniversitas int         `gorm:"null" json:"idUniversitas"`
	Universitas   Universitas `gorm:"foreignKey:IdUniversitas" json:"universitas"`

	//Register
	// NameRektor string `gorm:"size:255;not null;" json:"name_rektor"`
	// KtpRektor  string `gorm:"size:255;not null;" json:"ktp_rektor"`
	// Isverified bool   `gorm:"default:false" json:"isverified"`
	// Alamat     string `gorm:"null" json:"alamat"`
}
