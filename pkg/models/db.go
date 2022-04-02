package models

import (
	"github.com/shiowwj/go-cointracker-crud/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Transaction{})
}
