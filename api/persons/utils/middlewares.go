package utils

import (
	"fmt"

	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models"
	"github.com/gin-gonic/gin"
)

func GetUserAuth(c *gin.Context) (models.User, error) {
	sessionUser, exists := c.Get("authorizedUser")
	if !exists {
		return models.User{}, fmt.Errorf("the sessionUser was not found in the authorizedUser of the request context")
	}
	user := sessionUser.(models.User)
	return user, nil
}

func SetFiltersToUser(c *gin.Context, user *models.User) {
	queryParams := c.Request.URL.RawQuery
	user.SearchFilters = queryParams
}
