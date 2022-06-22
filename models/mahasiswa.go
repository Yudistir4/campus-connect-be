package models

import "gorm.io/gorm"

type Mahasiswa struct {
	gorm.Model
	Semester          uint
	Nim               string
	StatusMahasiswa   string
	IdUserUniversitas int
	Universitas       string

	// User              User `gorm:"foreignKey:IdUserUniversitas"`
	// idUserUniversitas: "foreignKey",
	// IdProdi: "foreignKey",
	// idFakultas: "foreignKey",
	// idJabatan: "foreignKey",
}
