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

type artilcleCreateForm struct {
	Article struct {
		Title       string   `form:"title" binding:"required"`
		Description string   `form:"description" binding:"required"`
		Body        string   `form:"Body" binding:"required"`
		TagList     []string `form:"tagLisg"`
	} `form:"article"`
}

type artilcleUpdateForm struct {
	Article struct {
		Title       string `form:"title"`
		Description string `form:"description"`
		Body        string `form:"Body"`
	} `form:"article"`
}

// GetArictles godoc
func GetArictles(c *gin.Context) {
	var query articlesQuery
	c.ShouldBindQuery(&query)

	var userID uint = 0
	if value, _ := c.Get(common.KeyJwtCurentUser); value != nil {
		var currentUserModel = value.(*userModels.UserModel)
		userID = currentUserModel.ID
	}

	articles, err := articleModels.FilterAirticles(query.Tag, query.Author, query.Favorited, query.Limit, query.Offset)
	if err != nil {
		panic(err)
	}

	articlesJSON := genArticlesData(articles, userID)

	c.JSON(http.StatusOK, gin.H{
		"articles":      articlesJSON,
		"articlesCount": len(articles),
	})
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

	articlesJSON := genArticlesData(articles, currentUserModel.ID)

	c.JSON(http.StatusOK, gin.H{"articles": articlesJSON, "articlesCount": len(articles)})
}

// GetArticle godoc
func GetArticle(c *gin.Context) {
	slug := c.Param("slug")

	articleModel, _ := articleModels.QueryArticle(slug)

	if articleModel.ID == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	articleSchema := genSingleArticleData(articleModel, 0)

	c.JSON(http.StatusOK, gin.H{"article": articleSchema})
}

// CreateArticle godoc
func CreateArticle(c *gin.Context) {
	var form artilcleCreateForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON("Title, Description, Body fields are required"))
		return
	}

	value, _ := c.Get(common.KeyJwtCurentUser)
	var currentUserModel = value.(*userModels.UserModel)

	articleModel, err := articleModels.AddArticle(
		form.Article.Title,
		form.Article.Description,
		form.Article.Body,
		currentUserModel.ID,
		form.Article.TagList,
	)

	if err != nil {
		panic(err)
	}

	articleSchema := genSingleArticleData(articleModel, currentUserModel.ID)

	c.JSON(http.StatusOK, gin.H{"article": articleSchema})
}

// UpdateArticle godoc
func UpdateArticle(c *gin.Context) {
	var form artilcleUpdateForm
	c.ShouldBindJSON(&form)
	originSlug := c.Param("slug")

	value, _ := c.Get(common.KeyJwtCurentUser)
	var currentUserModel = value.(*userModels.UserModel)

	if !articleModels.CheckArticleExistViaSlug(originSlug) {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON("article not exist"))
		return
	} else if article, _ := articleModels.QueryArticle(originSlug); article.AuthorID != currentUserModel.ID {
		c.AbortWithStatusJSON(http.StatusForbidden, common.GenErrorJSON("no permission modify the article"))
		return
	}

	articleModel, err := articleModels.ModifyArticle(
		originSlug,
		form.Article.Title,
		form.Article.Description,
		form.Article.Body,
	)

	if err != nil {
		panic(err)
	}

	articleSchema := genSingleArticleData(articleModel, currentUserModel.ID)

	c.JSON(http.StatusOK, gin.H{"article": articleSchema})
}

// RemoveArticle godoc
func RemoveArticle(c *gin.Context) {
	slug := c.Param("slug")

	value, _ := c.Get(common.KeyJwtCurentUser)
	var currentUserModel = value.(*userModels.UserModel)

	if !articleModels.CheckArticleExistViaSlug(slug) {
		// article not exist, directly return ok
		c.Status(http.StatusOK)
		return
	} else if article, _ := articleModels.QueryArticle(slug); article.AuthorID != currentUserModel.ID {
		c.AbortWithStatusJSON(http.StatusForbidden, common.GenErrorJSON("no permission modify the article"))
		return
	}

	if err := articleModels.RemoveArticle(slug); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

// FavoriteArticle godoc
func FavoriteArticle(c *gin.Context) {
	slug := c.Param("slug")

	value, _ := c.Get(common.KeyJwtCurentUser)
	var currentUserModel = value.(*userModels.UserModel)

	if articleModel, _ := articleModels.QueryArticle(slug); articleModel.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, common.GenErrorJSON("article not exist"))
	} else {
		err := articleModels.Favorite(currentUserModel.ID, articleModel.ID)
		if err != nil {
			panic(err)
		}

		// Requery for right data
		articleModel, _ := articleModels.QueryArticle(slug)
		c.JSON(http.StatusOK, gin.H{
			"article": genSingleArticleData(articleModel, currentUserModel.ID),
		})
	}
}

// UnFavoriteArticle godoc
func UnFavoriteArticle(c *gin.Context) {
	slug := c.Param("slug")

	value, _ := c.Get(common.KeyJwtCurentUser)
	var currentUserModel = value.(*userModels.UserModel)

	if articleModel, _ := articleModels.QueryArticle(slug); articleModel.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, common.GenErrorJSON("article not exist"))
	} else {
		err := articleModels.UnFavorite(currentUserModel.ID, articleModel.ID)
		if err != nil {
			panic(err)
		}

		// Requery for right data
		articleModel, _ := articleModels.QueryArticle(slug)

		c.JSON(http.StatusOK, gin.H{
			"article": genSingleArticleData(articleModel, currentUserModel.ID),
		})
	}
}

func genSingleArticleData(article *articleModels.ArticleModel, curUserID uint) *ArticleSchema {
	tagNames, _ := articleModels.GetArticleTagNames(article.ID)
	author, _ := userModels.GetUserByID(article.AuthorID)
	isFavorite := articleModels.GetFavoriteState(curUserID, article.ID)
	favoriteCount := articleModels.ArticleFavoritedCount(article.ID)

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
		Favorited:      isFavorite,
		FavoritesCount: favoriteCount,
	}

	return &articleSchema
}

func genArticlesData(articles []*articleModels.ArticleModel, curUserID uint) []*ArticleSchema {
	var articlesJSON = make([]*ArticleSchema, 0)

	for _, article := range articles {
		articlesJSON = append(articlesJSON, genSingleArticleData(article, curUserID))
	}

	return articlesJSON
}
