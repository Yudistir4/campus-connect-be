package models

import "time"

type Save struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	IdPost    int       `gorm:"not null" json:"idPost"`
	IdUser    int       `gorm:"not null" json:"idUser"`
	Post      Post      `gorm:"foreignKey:IdPost" json:"post"`
}
