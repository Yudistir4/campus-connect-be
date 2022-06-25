package models

import "time"

type Post struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	IdUser    int       `json:"idUser"`
	User      User      `gorm:"foreignKey:IdUser" json:"user"`

	Materi            string `json:"materi"`
	Caption           string `json:"caption"`
	IdUserUniversitas int    `json:"idUserUniversitas"` //for filtering
	IsNews            bool   `json:"isNews"`
}
