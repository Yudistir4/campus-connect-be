package models

import "time"

type Organisasi struct {
	ID                int       `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time `json:"createdAt"`
	IdUserUniversitas int       `gorm:"not null" json:"idUserUniversitas"`
	Universitas       string    `gorm:"not null" json:"universitas"`
}
