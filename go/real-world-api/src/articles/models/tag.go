package models

import (
	"real-world-api/src/db"
)

// TagModel godoc
type TagModel struct {
	ArticleID uint `gorm:"NOT NULL"`
	Name string `gorm:"NOT NULL;PRIMARY_KEY"`
}

// TableName godoc
func (TagModel) TableName() string {
	return "tag"
}

// GetTagRelatedArticles godoc
func GetTagRelatedArticles(tagName string) ([]uint, error) {
	var ids []uint
	err := db.GetDB().Find(TagModel{Name: tagName}).Pluck("article_id", &ids).Error
	return ids, err
}

// GetArticleTagNames godoc
func GetArticleTagNames(articleid uint) ([]string, error) {
	var tagNames []string = make([]string, 0)
	err := db.GetDB().Model(TagModel{}).Where("article_id=?", articleid).Pluck("name", &tagNames).Error
	return tagNames, err
}
