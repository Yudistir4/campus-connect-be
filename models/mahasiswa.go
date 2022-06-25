package models

import "time"

type Mahasiswa struct {
	ID                int       `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time `json:"createdAt"`
	Semester          uint      `gorm:"null" json:"semester"`
	Nim               string    `gorm:"not null" json:"nim"`
	StatusMahasiswa   string    `gorm:"not null" json:"statusMahasiswa"`
	IdUserUniversitas int       `gorm:"not null" json:"idUserUniversitas"`
	Universitas       string    `gorm:"not null" json:"universitas"`

	IdFakultas int      `gorm:"null" json:"idFakultas"`
	Fakultas   Fakultas `gorm:"foreignKey:IdFakultas" json:"fakultas"`

	IdProdi int   `gorm:"null" json:"idProdi"`
	Prodi   Prodi `gorm:"foreignKey:IdProdi" json:"prodi"`

	IdJabatan int     `gorm:"null" json:"idJabatan"`
	Jabatan   Jabatan `gorm:"foreignKey:IdJabatan" json:"jabatan"`
}
