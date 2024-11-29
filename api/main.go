package main

import (
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/database"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/shared"
)

func main() {

	// if os.Getenv("GIN_MODE") != "release" {
	// 	shared.LoadEnvVars()
	// }

	// shared.SetGinMode()
	database.CreateDbConnection()
	shared.MigrateDatabase()

	router := shared.SetupRouter()
	router.Run(":8080")
}
