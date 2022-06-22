package models

import "gorm.io/gorm"

type Universitas struct {
	gorm.Model
	NamaRektor string
	KtpRektor  string
	IsVerified bool
	Alamat     string
}
