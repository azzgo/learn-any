package models

import (
	"real-world-api/src/db"

	userModels "real-world-api/src/users/models"

	"github.com/jinzhu/gorm"
)

// ArticleModel godoc
type ArticleModel struct {
	gorm.Model
	Slug        string `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"size:512"`
	Body        string `gorm:"type:text"`
	AuthorID    uint
	// Favorited
	FavoritesCount uint `gorm:"default:0"`
}

// TableName godoc
func (ArticleModel) TableName() string {
	return "article"
}

// GetAirticles godoc
func GetAirticles(tag string, author string, favorited string, limit uint, offset uint) ([]*ArticleModel, error) {
	var articles = make([]*ArticleModel, 0)
	var ids []uint = make([]uint, 0)
	var err error

	if tag != "" {
		_ids, _ := GetTagRelatedArticles(tag)
		ids = append(ids, _ids...)
	}

	if author != "" {
		// Get AuthorID
		userModel, _ := userModels.GetUserByUsername(author)
		_ids, _ := GetAuthorRelatedAirticles(userModel.ID)
		ids = append(ids, _ids...)
	}

	if tag != "" || author != "" {
		err = db.GetDB().Where("id in (?)", ids).Offset(offset).Limit(limit).Order("updated_at desc").Find(&articles).Error	
	} else {
		err = db.GetDB().Offset(offset).Limit(limit).Order("updated_at desc").Find(&articles).Error	
	}

	return articles, err
}
