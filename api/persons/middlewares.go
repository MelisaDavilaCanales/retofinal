package persons

import (
	"net/http"

	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/database"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/persons/utils"
	"github.com/gin-gonic/gin"
)

func updateUserWithFilters() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.GetUserAuth(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "The sessionUser was not found in the authorizedUser of the request context."})
		}
		utils.SetFiltersToUser(c, &user)
		database.DBConn.Save(&user)
	}
}
