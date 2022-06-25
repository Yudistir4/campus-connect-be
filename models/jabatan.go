package models

import "time"

type Jabatan struct {
	ID               int       `gorm:"primaryKey" json:"id"`
	CreatedAt        time.Time `json:"createdAt"`
	NamaJabatan      string    `gorm:"not null" json:"namaJabatan"`
	NamaOrganisasi   string    `gorm:"not null" json:"namaOrganisasi"`
	IdMahasiswa      int       `json:"idMahasiswa"`
	IdUserMahasiswa  int       `gorm:"not null" json:"idUserMahasiswa"`
	IdUserOrganisasi int       `gorm:"not null" json:"idUserOrganisasi"`
	// User             User `gorm:"foreignKey:IdUserUniversitas"`
}
