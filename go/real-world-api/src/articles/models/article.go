package models

import (
	"github.com/jinzhu/gorm"
)


// ArticleModel godoc
type ArticleModel struct {
	gorm.Model
	slug string `gorm:"not null"`
	title string `gorm:"not null"`
	description string `gorm:"size:512"`
	body string `gorm:"type:text"`
	tagList []TagModel
	// favorited
	favoritesCount uint `gorm:"default:0"`
}