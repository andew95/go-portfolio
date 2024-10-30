package db

import (
	"fmt"
	"log"
	"portfolio/internal/repository/uploadRepository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(baseUrl string, username string, password string, dbName string) *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", username, password, baseUrl, dbName)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&uploadRepository.Upload{})

	return db
}
