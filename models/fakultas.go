package models

import "time"

type Fakultas struct {
	ID                int       `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time `json:"createdAt"`
	NamaFakultas      string    `gorm:"not null" json:"namaFakultas"`
	IdUserUniversitas int       `gorm:"not null" json:"idUserUniversitas"`
}
