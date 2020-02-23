package models

import (
	"real-world-api/src/db"

	"github.com/jinzhu/gorm"
)

// CommentModel godoc
type CommentModel struct {
	gorm.Model
	ArticleID uint `gorm:"not null;INDEX"`
	AuthorID uint `gorm:"not null;index"`
	Content string `gorm:"type:text"`
}

// TableName godoc
func (CommentModel) TableName() string {
	return "comment"
}

// AddComment godoc
func AddComment(articleSlug string, commentContent string, authorID uint) (*CommentModel, error) {

	articleModel, err := QueryArticle(articleSlug)
	if err != nil {
		return nil, err
	}

	commentModel := &CommentModel{
		ArticleID: articleModel.ID,
		AuthorID: authorID,
		Content: commentContent,
	}
	err = db.GetDB().Save(commentModel).Error

	return commentModel, err
}

