package statistics

import (
	"fmt"
	"net/http"

	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/database"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/persons/utils"
	"github.com/gin-gonic/gin"
)

type OccupationStats struct {
	Option     string `json:"option"`
	Percentage string `json:"percentage"`
}

func GetStatics(c *gin.Context, field string, valuesList []string) {
	var persons []models.Person
	var totalIncomeCount int64

	income := c.Query("income")

	query := database.DBConn.Model(&models.Person{})
	utils.AddFiltersToQuery(c, query, "income", income)
	query.Find(&persons).Count(&totalIncomeCount)

	valuesCounts := make(map[string]int64)
	for _, value := range valuesList {
		var count int64
		query := database.DBConn.Model(&models.Person{})
		utils.AddFiltersToQuery(c, query, "income", income)
		utils.AddFiltersToQuery(c, query, field, value)
		query.Find(&persons).Count(&count)
		valuesCounts[value] = count
	}

	var stats []OccupationStats
	for option, count := range valuesCounts {
		percentage := float64(count) / float64(totalIncomeCount) * 100
		stats = append(stats, OccupationStats{
			Option:     option,
			Percentage: fmt.Sprintf("%d", int(percentage)),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"statistics": stats,
	})
}
