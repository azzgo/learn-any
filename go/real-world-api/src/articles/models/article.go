package models

import (
	"real-world-api/src/db"

	userModels "real-world-api/src/users/models"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

// ArticleModel godoc
type ArticleModel struct {
	gorm.Model
	Slug        string `gorm:"not null;UNIQUE_INDEX"`
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

// FilterAirticles godoc
func FilterAirticles(tag string, author string, favorited string, limit uint, offset uint) ([]*ArticleModel, error) {
	return filterAirticlesQuery(nil, tag, author, favorited, limit, offset)
}

// FilterFollowingAuthorAirticles godoc
func FilterFollowingAuthorAirticles(curUserID uint, tag string, author string, favorited string, limit uint, offset uint) ([]*ArticleModel, error) {
	ids, err := userModels.GetTargetFollowingIDs("", curUserID)
	if err != nil {
		return nil, err
	}
	return filterAirticlesQuery(db.GetDB().Where("author_id in (?)", ids), tag, author, favorited, limit, offset)
}

// QueryArticle godoc
func QueryArticle(slug string) (*ArticleModel, error) {
	articleModel := new(ArticleModel)
	err := db.GetDB().Where("slug=?", slug).First(articleModel).Error
	return articleModel, err
}

// SaveArticle godoc
func SaveArticle(title string, description string, body string, authorID uint, tagList []string) (*ArticleModel, error) {
	slug := slug.Make(title)
	var articleModel *ArticleModel = new(ArticleModel)
	// if there one in database
	db.GetDB().Where("slug=?", slug).Find(articleModel)
	if articleModel.ID == 0 {
		articleModel = &ArticleModel{
			Title:       title,
			Slug:        slug,
			Description: description,
			Body:        body,
			AuthorID:    authorID,
		}
	}
	// Save article
	if err := db.GetDB().Save(articleModel).Error; err != nil {
		return nil, err
	}

	// Save related  tagsMapping
	if err := db.GetDB().Where("article_id = ?", articleModel.ID).Delete(TagModel{}).Error; err != nil {
		return nil, err
	} else if tagList != nil {
		err := db.GetDB().Transaction(func(tx *gorm.DB) error {
			for _, tagName := range(tagList) {
				if err := tx.Save(&TagModel{ArticleID: articleModel.ID, Name: tagName}).Error; err != nil {
					return err
				}
			}

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	// save related author mapping
	if err := db.GetDB().Save(&AuthorModel{AuthorID: authorID, ArticleID: articleModel.ID}).Error; err != nil {
		return nil, err
	}

	return articleModel, nil
}

func filterAirticlesQuery(preQuery *gorm.DB, tag string, author string, favorited string, limit uint, offset uint) ([]*ArticleModel, error) {
	var articles = make([]*ArticleModel, 0)
	var ids []uint = make([]uint, 0)
	var err error

	// judge use which scope
	var conn *gorm.DB
	if preQuery != nil {
		conn = preQuery
	} else {
		conn = db.GetDB()
	}

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
		err = conn.Where("id in (?)", ids).Offset(offset).Limit(limit).Order("updated_at desc").Find(&articles).Error
	} else {
		err = conn.Offset(offset).Limit(limit).Order("updated_at desc").Find(&articles).Error
	}

	return articles, err
}
