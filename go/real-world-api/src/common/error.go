package common

import "github.com/gin-gonic/gin"


// GenErrorJSON godoc
func GenErrorJSON(errorMessage string) gin.H {
	return gin.H{"errors": map[string][]string {
		"body": []string { errorMessage },
	}}
}

// GenErrorsJSON godoc
func GenErrorsJSON(errorMessages []string) gin.H {
	return gin.H{"errors": map[string][]string {
		"body": errorMessages,
	}}
}
