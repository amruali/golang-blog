package db

import (
	"github.com/amrali/golang-blog/pkg/models"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{},  &models.Post{}, &models.PostComment{} )
}