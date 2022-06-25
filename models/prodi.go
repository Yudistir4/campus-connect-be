package models

import "time"

type Prodi struct {
	ID                int       `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time `json:"createdAt"`
	NamaProdi         string    `gorm:"not null" json:"namaProdi"`
	IdUserUniversitas int       `gorm:"not null" json:"idUserUniversitas"`
	IdFakultas        int       `gorm:"not null" json:"idFakultas"`
	Fakultas          Fakultas  `gorm:"foreignKey:IdFakultas" json:"fakultas"`
	// IdUserUniversitas int
	// Universitas       string

	// User              User `gorm:"foreignKey:IdUserUniversitas"`
	// idUserUniversitas: "foreignKey",
	// IdProdi: "foreignKey",
	// idFakultas: "foreignKey",
	// idJabatan: "foreignKey",
}
