package models

import "time"

type Post struct {
	ID       uint           	`gorm:"AUTO_INCREMENT"`
	Description  string         `gorm:"type:VARCHAR(400);NOT NULL"`
	UserID       uint           
	PostComment  []PostComment 
	CreatedAt    time.Time      
	UpdatedAt    time.Time      
}