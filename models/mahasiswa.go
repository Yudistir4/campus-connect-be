package models

import "gorm.io/gorm"

type Mahasiswa struct {
	gorm.Model
	Semester          uint
	Nim               string
	StatusMahasiswa   string
	IdUserUniversitas int
	Universitas       string

	IdFakultas int
	Fakultas   Fakultas `gorm:"foreignKey:IdFakultas"`
	IdProdi    int
	Prodi      Prodi `gorm:"foreignKey:IdProdi"`
	IdJabatan  int
	Jabatan    Jabatan `gorm:"foreignKey:IdJabatan"`
}
