package models

import (
	"real-world-api/src/db"
)

// FavoriteModel godoc
type FavoriteModel struct {
	ArticleID uint `gorm:"NOT NULL;PRIMARY_KEY"`
	UserID    uint `gorm:"NOT NULL;UNIQUE_INDEX"`
}

// TableName godoc
func (FavoriteModel) TableName() string {
	return "favorite"
}

// ArticleFavoritedCount godoc
func ArticleFavoritedCount(articleID uint) uint {
	var count uint = 0
	db.GetDB().Model(FavoriteModel{}).Where("article_id=?", articleID).Count(&count)
	return count
}

// GetFavoriteState godoc
func GetFavoriteState(userID uint, articleID uint) bool {
	var count uint = 0
	db.GetDB().Model(FavoriteModel{}).Where("article_id=? && user_id=?", articleID, userID).Count(&count)
	return count == 1
}

// Favorite godoc
func Favorite(userID uint, articleID uint) error {
	return db.GetDB().Save(&FavoriteModel{UserID: userID, ArticleID: articleID}).Error
}

// UnFavorite godoc
func UnFavorite(userID uint, articleID uint) error {
	return db.GetDB().Delete(&FavoriteModel{UserID: userID, ArticleID: articleID}).Error
}
