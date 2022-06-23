package models

import "gorm.io/gorm"

type Fakultas struct {
	gorm.Model
	NamaFakultas      string
	IdUserUniversitas int
	// IdUserUniversitas int
	// Universitas       string

	// User              User `gorm:"foreignKey:IdUserUniversitas"`
	// idUserUniversitas: "foreignKey",
	// IdProdi: "foreignKey",
	// idFakultas: "foreignKey",
	// idJabatan: "foreignKey",
}
