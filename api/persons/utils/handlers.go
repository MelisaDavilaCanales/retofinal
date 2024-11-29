package utils

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageStr := c.DefaultQuery("page", "1")
		page, _ := strconv.Atoi(pageStr)
		if page <= 0 {
			page = 1
		}

		pageSizeStr := c.DefaultQuery("page_size", "10")
		pageSize, _ := strconv.Atoi(pageSizeStr)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		dbClone := db.Session(&gorm.Session{})
		var total int64
		dbClone.Count(&total)

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

		c.Set("page", page)
		c.Set("pageSize", pageSize)
		c.Set("totalPages", totalPages)

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func DecodeAndSplit(param string) ([]string, error) {
	if param == "" {
		return nil, nil
	}
	decodedParam, err := url.QueryUnescape(param)
	if err != nil {
		return nil, err
	}

	items := strings.Split(decodedParam, ",")
	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}

	return items, nil
}

func AddFiltersToQuery(c *gin.Context, query *gorm.DB, field string, values string) error {
	valueList, err := DecodeAndSplit(values)
	if err != nil || len(valueList) == 0 {
		c.JSON(400, gin.H{"invalid parameter": field})
		return fmt.Errorf("invalid %s parameter", field)
	}
	*query = *query.Where(fmt.Sprintf("%s IN ?", field), valueList)
	return nil
}

// Returns a map with the field and its values for each query parameter
func GetINQueryParameters(c *gin.Context) map[string]string {
	education := c.Query("education") //"Bachelors,Professional school"
	maritalStatus := c.Query("marital_status")
	occupation := c.Query("occupation")
	income := c.Query("income")

	parameterList := make(map[string]string)

	parameterList["education"] = education
	parameterList["marital_status"] = maritalStatus
	parameterList["occupation"] = occupation
	parameterList["income"] = income

	return parameterList
}

func GetBETWEENQueryParameters(c *gin.Context) map[string]string {
	ageRange := c.Query("age_range")

	parameterList := make(map[string]string)
	parameterList["age"] = ageRange

	return parameterList
}

func AplyINFilters(c *gin.Context, query *gorm.DB, parameterList map[string]string) error {
	for field, values := range parameterList {
		if values != "" {
			err := AddFiltersToQuery(c, query, field, values)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func AplyBETWEENFilters(c *gin.Context, query *gorm.DB, parameterList map[string]string) error {
	for field, values := range parameterList {
		if values != "" {
			ageRangesList, err := DecodeAndSplit(values)
			if err != nil {
				c.JSON(400, gin.H{"Error Invalid parameter": field})
				return err
			}
			for _, r := range ageRangesList {
				parts := strings.Split(r, "-")
				if len(parts) == 2 {
					lowerBound, _ := strconv.Atoi(parts[0])
					upperBound, _ := strconv.Atoi(parts[1])

					*query = *query.Or(fmt.Sprintf("%s BETWEEN %d AND %d", field, lowerBound, upperBound))
				}
			}
		}
	}
	return nil
}
