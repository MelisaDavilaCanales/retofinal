package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func CreateDbConnection() (*gorm.DB, error) {
	// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )
	connStr := "host=db port=5432 user=me password=supersecret dbname=trucode sslmode=disable"
	var err error
	DBConn, err = gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})
	if err != nil || DBConn == nil {
		log.Fatal("Failed to connect to the database", err)
	}
	return DBConn, err
}
