package statistics

import (
	"github.com/gin-gonic/gin"
)

func GetOccupationStatistics(c *gin.Context) {
	field := "occupation"
	occupationList := []string{
		"Tech-support", "Craft-repair", "Other-service", "Sales", "Exec-managerial", "Prof-specialty",
		"Handlers-cleaners", "Machine-op-inspct", "Adm-clerical", "Farming-fishing", "Transport-moving",
		"Priv-house-serv", "Protective-serv", "Armed-Forces",
	}

	GetStatics(c, field, occupationList)
}

func GetEducationStatistics(c *gin.Context) {
	field := "education"
	educationList := []string{
		"Bachelors", "Some-college", "11th", "HS-grad", "Prof-school", "Assoc-acdm",
		"Assoc-voc", "9th", "7th-8th", "12th", "Masters", "1st-4th", "10th",
		"Doctorate", "5th-6th", "Preschool",
	}

	GetStatics(c, field, educationList)
}

func GetCountryStatistics(c *gin.Context) {
	field := "native_country"
	countryList := []string{
		"United-States", "Cambodia", "England", "Puerto-Rico", "Canada", "Germany",
		"Outlying-US(Guam-USVI-etc)", "India", "Japan", "Greece", "South", "China",
		"Cuba", "Iran", "Honduras", "Philippines", "Italy", "Poland", "Jamaica",
		"Vietnam", "Mexico", "Portugal", "Ireland", "France", "Dominican-Republic",
		"Laos", "Ecuador", "Taiwan", "Haiti", "Columbia", "Hungary", "Guatemala",
		"Nicaragua", "Scotland", "Thailand", "Yugoslavia", "El-Salvador",
		"Trinidad&Tobago", "Peru", "Hong", "Holand-Netherlands",
	}

	GetStatics(c, field, countryList)
}
func GetGenderStatistics(c *gin.Context) {
	field := "sex"
	countryList := []string{
		"Female", "Male",
	}

	GetStatics(c, field, countryList)
}
