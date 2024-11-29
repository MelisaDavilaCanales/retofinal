package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	sessionUser, _ := c.Get("authorizedUser")
	c.JSON(http.StatusOK, sessionUser)
}
