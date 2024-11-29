package utils

import (
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/database"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/persons/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApplyFilters(c *gin.Context) *gorm.DB {
	query := database.DBConn.Model(&models.Person{})

	INParameterList := utils.GetINQueryParameters(c)
	utils.AplyINFilters(c, query, INParameterList)

	BETWEENParameters := utils.GetBETWEENQueryParameters(c)
	utils.AplyBETWEENFilters(c, query, BETWEENParameters)

	return query
}
