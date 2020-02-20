package handlers

import (
	"net/http"

	articleModels "real-world-api/src/articles/models"
	userHanlders "real-world-api/src/users/handlers"
	userModels "real-world-api/src/users/models"

	"github.com/gin-gonic/gin"
)

// ArictlesQuery godoc
type arictlesQuery struct {
	Tag       string `form:"tag"`
	Author    string `form:"Author"`
	Favorited string `form:"favorited"`
	Limit     uint   `form:"limit,default=20"`
	Offset    uint   `form:"offset"`
}

// GetArictles godoc
func GetArictles(c *gin.Context) {
	var query arictlesQuery
	c.ShouldBindQuery(&query)

	articles, err := articleModels.GetAirticles(query.Tag, query.Author, query.Favorited, query.Limit, query.Offset)
	if err != nil {
		panic(err)
	}

	articlesJSON := genArticlesData(articles)

	c.JSON(http.StatusOK, gin.H{"articles": articlesJSON})
}

// GetArticlesByFeed godoc
func GetArticlesByFeed(c *gin.Context) {
	var query arictlesQuery
	c.ShouldBindQuery(&query)

	articles, err := articleModels.GetAirticles(query.Tag, query.Author, query.Favorited, query.Limit, query.Offset)
	if err != nil {
		panic(err)
	}

	articlesJSON := genArticlesData(articles)

	c.JSON(http.StatusOK, gin.H{"articles": articlesJSON})
}

func genArticlesData(articles []*articleModels.ArticleModel) []ArticleSchema {
	var articlesJSON = make([]ArticleSchema, 0)

	for _, article := range articles {
		tagNames, _ := articleModels.GetArticleTagNames(article.ID)
		author, _ := userModels.GetUserByID(article.AuthorID)

		articlesJSON = append(articlesJSON, ArticleSchema{
			TagList: tagNames,
			Author: userHanlders.ProfileSchema{
				Username:  author.Username,
				Image:     author.Image,
				Following: false,
				Bio:       author.Bio,
			},
			Slug:           article.Slug,
			Title:          article.Title,
			Description:    article.Description,
			Body:           article.Body,
			CreateAt:       article.CreatedAt,
			UpdateAt:       article.UpdatedAt,
			Favorited:      false,
			FavoritesCount: article.FavoritesCount,
		})
	}

	return articlesJSON
}
