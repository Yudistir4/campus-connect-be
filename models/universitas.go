package models

import "time"

type Universitas struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	NamaRektor string    `gorm:"not null" json:"namaRektor"`
	KtpRektor  string    `gorm:"not null" json:"ktpRektor"`
	IsVerified bool      `gorm:"not null" json:"isVerified"`
	Alamat     string    `gorm:"not null" json:"alamat"`
}
