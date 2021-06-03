package models

import "time"

type PostComment struct {
	ID          uint   `gorm:"AUTO_INCREMENT"`
	Description string `gorm:"type:VARCHAR(400);NOT NULL"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uint
	PostID      uint
}
