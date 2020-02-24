package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	articleModels "real-world-api/src/articles/models"
)

// QueryTagList godoc
func QueryTagList(c *gin.Context)  {
	if tags, err := articleModels.GetTagList(); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"tags": tags,
		})
	}
}