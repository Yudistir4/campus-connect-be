package models

import "time"

type Comment struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Comment   string    `gorm:"not null" json:"comment"`
	IdPost    int       `gorm:"not null" json:"idPost"`
	IdUser    int       `gorm:"not null" json:"idUser"`
	User      User      `gorm:"foreignKey:IdUser" json:"user"`
}
