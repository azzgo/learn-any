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

// QueryComments godoc
func QueryComments(slug string) ([]*CommentModel, error) {
	articleModel, err := QueryArticle(slug)

	if err != nil {
		return nil, err
	}

	var result []*CommentModel = make([]*CommentModel, 0)

	err = db.GetDB().Model(CommentModel{}).Where("article_id=?", articleModel.ID).Scan(&result).Error

	return result, err
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


// RemoveComment godoc
func RemoveComment(slug string, commentID uint) error {
	articleModel, err := QueryArticle(slug)
	if err != nil {
		return err
	}

	err = db.GetDB().Delete(&CommentModel{
		ArticleID: articleModel.ID,
		Model: gorm.Model {
			ID: commentID,
		},
	}).Error

	return err
}
