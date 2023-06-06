package configuration

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config Config) (readDB, writeDB *gorm.DB, err error) {
	// Construct the DSN string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.Get("DB_HOST"), config.Get("DB_PORT"), config.Get("DB_USER"), config.Get("DB_PASS"), config.Get("DB_NAME"),
	)

	// Open new connection to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	readDB = db
	writeDB = db
	return
}
