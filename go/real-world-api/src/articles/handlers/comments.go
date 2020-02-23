package handlers

import (
	"net/http"
	"real-world-api/src/common"
	userHanlders "real-world-api/src/users/handlers"
	userModels "real-world-api/src/users/models"

	articleModels "real-world-api/src/articles/models"

	"github.com/gin-gonic/gin"
	"strconv"
)

type commentCreateForm struct {
	Comment struct {
		Body string `form:"body" binding:"required"`
	} `form:"comment"`
}

func QueryComments(c *gin.Context) {
	slug := c.Param("slug")

	if results, err := articleModels.QueryComments(slug); err != nil {
		panic(err)
	} else {
		commentsResult := make([]*CommentSchema, 0)
		for _, article := range results {
			commentsResult = append(commentsResult, genSingleCommentSchema(article))
		}
		c.JSON(http.StatusOK, gin.H {
			"comments": commentsResult,
		})
	}
}

// AddComent godoc
func AddComent(c *gin.Context) {
	var form commentCreateForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.GenErrorJSON("body is required"))
	}

	value, _ := c.Get(common.KeyJwtCurentUser)
	var currentUserModel = value.(*userModels.UserModel)

	slug := c.Param("slug")

	commentModel, err := articleModels.AddComment(slug, form.Comment.Body, currentUserModel.ID)
	if err != nil {
		panic(err)
	}

	commentSchema := genSingleCommentSchema(commentModel)

	c.JSON(http.StatusOK, gin.H{"comment": commentSchema})
}

func genSingleCommentSchema(commentModel *articleModels.CommentModel) *CommentSchema {
	author, _ := userModels.GetUserByID(commentModel.AuthorID)

	commentSchema := &CommentSchema{
		ID:        commentModel.ID,
		CreatedAt: commentModel.CreatedAt,
		UpdatedAt: commentModel.UpdatedAt,
		Body:      commentModel.Content,
		Author: userHanlders.ProfileSchema{
			Username:  author.Username,
			Image:     author.Image,
			Following: false,
			Bio:       author.Bio,
		},
	}
	return commentSchema
}


// RemoveComment godoc
func RemoveComment(c *gin.Context)  {
	slug := c.Param("slug")
	commentID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	err := articleModels.RemoveComment(slug, uint(commentID))
	if err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
	return
}