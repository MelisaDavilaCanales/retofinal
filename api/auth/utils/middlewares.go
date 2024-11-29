package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	tokenSlice := strings.Split(bearerToken, " ")
	if len(tokenSlice) == 2 {
		return tokenSlice[1]
	}
	return ""
}
