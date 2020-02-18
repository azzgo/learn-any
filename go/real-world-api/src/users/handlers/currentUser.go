package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CurrentUser godoc
// @tags Users
// @Accept  json
// @Produce  json
// @Produce  json
// @Router /user [get]
func CurrentUser(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message": "in building"})
}