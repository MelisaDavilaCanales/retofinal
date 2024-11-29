package persons

import (
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models"
	personsUtils "github.com/MelisaDavilaCanales/trucode3-challenge-final.git/persons/utils"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/utils"
	"github.com/gin-gonic/gin"
)

func GetPersons(c *gin.Context) {
	var persons []models.Person

	query := utils.ApplyFilters(c)
	result := query.Find(&persons)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error Failed to retrieve persons with filters from the database": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"persons": persons,
	})
}

func GetPaginatedPersons(c *gin.Context) {
	var persons []models.Person
	query := utils.ApplyFilters(c)
	query.Scopes(personsUtils.Paginate(c))
	result := query.Find(&persons)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error Failed to retrieve persons with filters from the database": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"persons":    persons,
		"page":       c.GetInt("page"),
		"pageSize":   c.GetInt("pageSize"),
		"totalPages": c.GetInt("totalPages"),
	})
}
