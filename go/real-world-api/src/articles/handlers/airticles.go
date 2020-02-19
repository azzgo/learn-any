package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type arictlesQuery struct {
	Tag string `form:"tag"`
	Author string `form:"Author"`
	Favorited string `form:"favorited"`
	Limit uint `form:"limit,default=20"`
	Offset uint `form:"offset"`
}

// GetArictles godoc
func GetArictles(c *gin.Context)  {
	var query arictlesQuery
	c.ShouldBindQuery(&query)

	c.JSON(http.StatusOK, gin.H {"status": "in building"})
}
