package models

import "gorm.io/gorm"

type Jabatan struct {
	gorm.Model
	NamaJabatan      string
	NamaOrganisasi   string
	IdUserMahasiswa  int
	IdUserOrganisasi int
	// User             User `gorm:"foreignKey:IdUserUniversitas"`
}
