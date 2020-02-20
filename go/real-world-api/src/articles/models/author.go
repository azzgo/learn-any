package models

import (
	"real-world-api/src/db"
)

// AuthorModel godoc
type AuthorModel struct {
	ArticleID uint `gorm:"NOT NULL"`
	AuthorID uint `gorm:"NOT NULL;PRIMARY_KEY"`
}

// TableName godoc
func (AuthorModel) TableName() string {
	return "author"
}

// GetAuthorRelatedAirticles godoc
func GetAuthorRelatedAirticles(authorid uint) ([]uint, error) {
	var airticles []uint
	err := db.GetDB().Model(AuthorModel{AuthorID: authorid}).Find(&airticles).Error
	return airticles, err
}