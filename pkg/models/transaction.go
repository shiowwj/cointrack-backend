package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey" json:"-"`
	Type         string    `json:"type"`
	Symbol       string    `json:"symbol"`
	Name         string    `json:"name"`
	Quantity     float64   `json:"quantity"`
	Total        float64   `json:"total"`
	PricePerCoin float64   `json:"pricePerCoin"`
	DateTimeTxn  time.Time `json:"dateTimeTxn"`
	Fees         float64   `json:"fees"`
	UserUUID     string    `json:"-"`
}

// `json:"userUuid"`

func (t *Transaction) CreateTransaction(userId string) *Transaction {
	t.UserUUID = strings.ToLower(userId)
	db.Create(&t)
	return t
}

func GetAllTransactions() []Transaction {
	var Transactions []Transaction
	db.Find(&Transactions)
	return Transactions
}

func GetTransactionById(id int64) (*Transaction, *gorm.DB) {
	var getTransaction Transaction
	// db := db.Where("ID=?", id).Find(&getHolding)
	db.First(&getTransaction, id)
	return &getTransaction, db
}

// func GetTransactionsBySymbol(symbol string) (*[]Transaction, *gorm.DB) {
// 	log.Info("GetTransactionsBySymbol", zap.String("symbol", symbol))
// 	transactions := make([]Transaction, 0)

// 	db.Where("symbol=?", symbol).Find(&transactions)

// 	return &transactions, db
// }

func DeleteTransaction(id int64) Transaction {
	var transaction Transaction
	db.Delete(&Transaction{}, id)
	return transaction
}

func getTransactionsByUserId(userId string) ([]AssetTransaction, *gorm.DB) {
	transactions := make([]Transaction, 0)
	db.Where("user_uuid=?", strings.ToLower(userId)).Find(&transactions)
	txns := make([]AssetTransaction, 0)
	for _, t := range transactions {
		_t := AssetTransaction{
			TxnId:        t.ID,
			Type:         t.Type,
			Symbol:       t.Symbol,
			Name:         t.Name,
			Quantity:     t.Quantity,
			Total:        t.Total,
			PricePerCoin: t.PricePerCoin,
			DateTimeTxn:  t.DateTimeTxn,
			Fees:         t.Fees,
		}
		txns = append(txns, _t)
	}

	return txns, db
}
