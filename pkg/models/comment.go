package models

type PostComment struct {
	ID          uint   `gorm:"AUTO_INCREMENT"`
	Description string `gorm:"type:VARCHAR(400);NOT NULL"`
	UserID      uint
	PostID      uint
}
