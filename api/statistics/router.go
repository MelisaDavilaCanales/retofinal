package statistics

import (
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/auth"
	"github.com/gin-gonic/gin"
)

func AddStaisticsRoute(c *gin.Engine) {
	persons := c.Group("/statistics")
	persons.Use(auth.AuthenticateSession())
	persons.GET("/ocupation", GetOccupationStatistics)
	persons.GET("/education", GetEducationStatistics)
	persons.GET("/country", GetCountryStatistics)
	persons.GET("/gender", GetGenderStatistics)
}
