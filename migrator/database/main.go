package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func CreateDbConnection() *gorm.DB {

	connStr := "host=db port=5432 user=me password=supersecret dbname=trucode sslmode=disable"
	// connStr := "host=localhost port=5432 user=me password=supersecret dbname=trucode sslmode=disable"
	var err error
	DBConn, err = gorm.Open(postgres.Open(connStr))
	if err != nil {
		log.Fatal("Error conexion con la BD")
	}
	return DBConn
}

// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"),
// 	os.Getenv("DB_PORT"),
// 	os.Getenv("DB_USER"),
// 	os.Getenv("DB_PASSWORD"),
// 	os.Getenv("DB_NAME"),
// 	os.Getenv("DB_SSLMODE"),
// )
