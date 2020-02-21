package handlers

import (
	"net/http"

	articleModels "real-world-api/src/articles/models"
	"real-world-api/src/common"
	userHanlders "real-world-api/src/users/handlers"
	userModels "real-world-api/src/users/models"

	"github.com/gin-gonic/gin"
)

// ArictlesQuery godoc
type articlesQuery struct {
	Tag       string `form:"tag"`
	Author    string `form:"Author"`
	Favorited string `form:"favorited"`
	Limit     uint   `form:"limit,default=20"`
	Offset    uint   `form:"offset"`
}

type artilcleForm struct {
	Article struct {
		Title string `form:"title" binding:"required"`
		Description string `form:"description" binding:"required"`
		Body string `form:"Body" binding:"required"`
		TagList []string `form:"tagLisg"`
	} `form:"article"`
}

// GetArictles godoc
func GetArictles(c *gin.Context) {
	var query articlesQuery
	c.ShouldBindQuery(&query)

	articles, err := articleModels.FilterAirticles(query.Tag, query.Author, query.Favorited, query.Limit, query.Offset)
	if err != nil {
		panic(err)
	}

	articlesJSON := genArticlesData(articles)

	c.JSON(http.StatusOK, gin.H{"articles": articlesJSON})
}

// GetArticlesByFeed godoc
func GetArticlesByFeed(c *gin.Context) {
	var query articlesQuery
	c.ShouldBindQuery(&query)

	value, _ := c.Get(common.KeyJwtCurentUser)
	var currentUserModel = value.(*userModels.UserModel)

	articles, err := articleModels.FilterFollowingAuthorAirticles(currentUserModel.ID, query.Tag, query.Author, query.Favorited, query.Limit, query.Offset)
	if err != nil {
		panic(err)
	}

	articlesJSON := genArticlesData(articles)

	c.JSON(http.StatusOK, gin.H{"articles": articlesJSON})
}

// GetArticle godoc
func GetArticle(c *gin.Context)  {
	slug := c.Param("slug")

	articleModel, _ := articleModels.QueryArticle(slug)

	articleSchema := genSingleArticleData(articleModel)
	
	c.JSON(http.StatusOK, gin.H{"article": articleSchema})
}

// CreateArticle godoc
func CreateArticle(c *gin.Context)  {
	var form artilcleForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON("Title, Description, Body fields are required"))
		return
	}

	value, _ := c.Get(common.KeyJwtCurentUser)
	var currentUserModel = value.(*userModels.UserModel)

	if exist := articleModels.CheckArticleExist(form.Article.Title); !exist {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON("article exists"))
		return
	}

	articleModel, err := articleModels.SaveArticle(
		form.Article.Title,
		form.Article.Description,
		form.Article.Body,
		currentUserModel.ID,
		form.Article.TagList,
	)

	if err != nil {
		panic(err)
	}

	articleSchema := genSingleArticleData(articleModel)


	c.JSON(http.StatusOK, gin.H {"article": articleSchema })
}

func genSingleArticleData(article *articleModels.ArticleModel) *ArticleSchema {
	tagNames, _ := articleModels.GetArticleTagNames(article.ID)
	author, _ := userModels.GetUserByID(article.AuthorID)

	articleSchema := ArticleSchema{
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
	}

	return &articleSchema
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
