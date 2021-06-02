package models

import "time"

type User struct {
	ID           uint   `gorm:"AUTO_INCREMENT"`
	UserName     string `gorm:"type:VARCHAR(50);UNIQUE;NOT NULL" json:"user_name"`
	Email        string `gorm:"type:VARCHAR(100);UNIQUE;NOT NULL" json:"email"`
	Password     []byte `json:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Posts        []Post
	PostComments []PostComment
}
