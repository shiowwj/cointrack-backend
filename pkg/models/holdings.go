package models

import (
	"log"

	"github.com/shiowwj/go-cointracker-crud/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Holding struct {
	gorm.Model
	ID           int64   `gorm:"primaryKey""json:"id"`
	Symbol       string  `json:"symbol"`
	AmountBought float64 `json:"amountBought"`
	Currency     string  `json:"currency"`

	BoughtAtMarketPrice float64 `json:"boughtAtMarketPrice"`
	// date format => string: "yyyy/mm/dd"
	BoughtAtDate string `json:"boughtAtDate"`

	SoldAtMarketPrice float64 `json:"soldAtMarketPrice"`
	// date format => string: "yyyy/mm/dd"
	SoldAtDate string `json:"soldAtDate"`

	Status string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Holding{})
}

func (h *Holding) CreateHolding() *Holding {
	db.Create(&h)
	return h
}

func GetAllHoldings() []Holding {
	var Holdings []Holding
	db.Find(&Holdings)
	return Holdings
}

func GetHoldingById(id int64) (*Holding, *gorm.DB) {
	var getHolding Holding
	// db := db.Where("ID=?", id).Find(&getHolding)
	db.First(&getHolding, id)
	return &getHolding, db
}

func GetHoldingsBySymbol(symbol string) (*[]Holding, *gorm.DB) {
	log.Println("symbol", symbol)
	holdings := make([]Holding, 0)

	db.Where("symbol=?", symbol).Find(&holdings)

	return &holdings, db
}

func DeleteHolding(id int64) Holding {
	var holding Holding
	db.Delete(&Holding{}, id)
	return holding
}
