package shared

import (
	"log"

	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/auth"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/database"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/persons"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/statistics"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func MigrateDatabase() {
	database.DBConn.AutoMigrate(&models.User{})
}

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// func SetGinMode() {
// 	appMode := os.Getenv("APP_MODE")
// 	if appMode == "production" {
// 		gin.SetMode(gin.ReleaseMode)
// 	}
// }

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	users.AddUserRoutes(router)
	auth.AddAuthRoutes(router)
	persons.AddersonRoutes(router)
	statistics.AddStaisticsRoute(router)

	return router
}
