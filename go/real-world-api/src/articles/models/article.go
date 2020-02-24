package models

import (
	"real-world-api/src/db"

	userModels "real-world-api/src/users/models"

	"github.com/google/uuid"
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
	if err := db.GetDB().Where("slug=?", slug).First(articleModel).Error; err != nil {
		return nil, err
	}

	return articleModel, nil
}

// CheckArticleExistViaTitle godoc
func CheckArticleExistViaTitle(title string) bool {
	var count uint
	db.GetDB().Model(ArticleModel{}).Where("title=?", title).Count(&count)
	return count == 1
}

// CheckArticleExistViaSlug godoc
func CheckArticleExistViaSlug(slug string) bool {
	var count uint
	db.GetDB().Model(ArticleModel{}).Where("slug=?", slug).Count(&count)
	return count == 1
}

// AddArticle godoc
func AddArticle(title string, description string, body string, authorID uint, tagList []string) (*ArticleModel, error) {
	slug := slug.Make(title)

	// if slug already exist
	if CheckArticleExistViaSlug(slug) {
		slug = slug + uuid.New().String()
	}

	articleModel := &ArticleModel{
		Title:       title,
		Slug:        slug,
		Description: description,
		Body:        body,
		AuthorID:    authorID,
	}
	// Save article
	if err := db.GetDB().Save(articleModel).Error; err != nil {
		return nil, err
	}

	// Add related  tagsMapping
	if tagList != nil {
		err := db.GetDB().Transaction(func(tx *gorm.DB) error {
			for _, tagName := range tagList {
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

// ModifyArticle godoc
func ModifyArticle(originSlug string, title string, description string, body string) (*ArticleModel, error) {
	articleModel, err := QueryArticle(originSlug)
	if err != nil {
		return nil, err
	}

	if title != "" {
		var newSlug string
		if title != articleModel.Title {
			newSlug = slug.Make(title)
			if CheckArticleExistViaTitle(title) {
				newSlug += uuid.New().String()
			}
			articleModel.Title = title
			articleModel.Slug = newSlug
		}
	} else {
		articleModel.Slug = originSlug
	}

	if description != "" {
		articleModel.Description = description
	}

	if body != "" {
		articleModel.Body = body
	}

	if err := db.GetDB().Save(articleModel).Error; err != nil {
		return nil, err
	}

	return articleModel, nil
}

// RemoveArticle godoc
func RemoveArticle(slug string) error {
	return db.GetDB().Delete(ArticleModel{}, "slug=?", slug).Error
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
