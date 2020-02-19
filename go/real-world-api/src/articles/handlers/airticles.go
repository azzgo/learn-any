package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetArictles godoc
func GetArictles(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H {"status": "in building"})
}
