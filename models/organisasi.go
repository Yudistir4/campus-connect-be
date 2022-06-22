package models

import "gorm.io/gorm"

type Organisasi struct {
	gorm.Model
	IdUserUniversitas int
	Universitas       string
	// User              User `gorm:"foreignKey:IdUserUniversitas"`
	// IdProdi: "foreignKey",
	// idFakultas: "foreignKey",
	// idUserUniversitas: "foreignKey",
	// idJabatan: "foreignKey",
}
