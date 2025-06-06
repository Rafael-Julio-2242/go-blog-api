package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(conString string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
