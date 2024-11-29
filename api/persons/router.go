package persons

import (
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/auth"
	"github.com/gin-gonic/gin"
)

func AddersonRoutes(c *gin.Engine) {
	persons := c.Group("/persons")
	persons.Use(auth.AuthenticateSession())
	persons.GET("", updateUserWithFilters(), GetPersons)
	persons.GET("/paginated", updateUserWithFilters(), GetPaginatedPersons)
}
