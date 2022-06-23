package models

import "gorm.io/gorm"

type Prodi struct {
	gorm.Model
	NamaProdi         string
	IdUserUniversitas int
	IdFakultas        int
	Fakultas          Fakultas `gorm:"foreignKey:IdFakultas"`
	// IdUserUniversitas int
	// Universitas       string

	// User              User `gorm:"foreignKey:IdUserUniversitas"`
	// idUserUniversitas: "foreignKey",
	// IdProdi: "foreignKey",
	// idFakultas: "foreignKey",
	// idJabatan: "foreignKey",
}
