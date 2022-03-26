package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// dsn := "host=localhost user=shwj password=1234 dbname=cointracker_test port=5423 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=localhost user=shwj password=1234 dbname=cointracker_test port=5432"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
