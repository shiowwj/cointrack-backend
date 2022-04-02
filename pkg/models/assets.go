package models

import (
	"strings"
	"time"

	"github.com/shiowwj/go-cointracker-crud/pkg/utils/log"
	"go.uber.org/zap"
)

type Portfolio struct {
	TotalAssetsBalance float64          `json:"totalAssetsBalance"`
	UserId             string           `json:"userid"`
	Assets             map[string]Asset `json:"assets"`
	AssetCount         int              `json:"-"`
}

type AssetTransaction struct {
	TxnId        uint      `json:"txnId"`
	Type         string    `json:"type"`
	Symbol       string    `json:"symbol"`
	Name         string    `json:"name"`
	Quantity     float64   `json:"quantity"`
	Total        float64   `json:"total"`
	PricePerCoin float64   `json:"pricePerCoin"`
	DateTimeTxn  time.Time `json:"dateTimeTxn"`
	Fees         float64   `json:"fees"`
}

type Asset struct {
	TotalValue   float64            `json:"totalValue"`
	TotalHolding float64            `json:"totalHolding"`
	AveragePrice float64            `json:"averagePrice"`
	Name         string             `json:"name"`
	Symbol       string             `json:"symbol"`
	Transactions []AssetTransaction `json:"transactions"`
}

func GetUserAssets(userId string) (Portfolio, error) {
	log.Info("GetUserAssets", zap.String("userId", userId))
	var portfolio Portfolio
	portfolio.Assets = make(map[string]Asset)
	portfolio.UserId = userId
	// Get all transcations for userId
	txns, db := getTransactionsByUserId(strings.ToLower(userId))
	if db.Error != nil {
		return portfolio, db.Error
	}
	for _, txn := range txns {
		if a, found := portfolio.Assets[txn.Symbol]; !found {
			// var asset Assetass
			asset := Asset{
				Name:   txn.Name,
				Symbol: txn.Symbol,
			}
			asset.Transactions = append(asset.Transactions, txn)
			portfolio.Assets[txn.Symbol] = asset
		} else {
			a.Transactions = append(a.Transactions, txn)
			portfolio.Assets[txn.Symbol] = a
		}
	}

	for sym, asset := range portfolio.Assets {
		var totalVal, totalHoldings float64
		for _, t := range asset.Transactions {
			totalVal = totalVal + t.Total
			totalHoldings = totalHoldings + t.Quantity
		}
		asset.TotalValue = totalVal
		asset.TotalHolding = totalHoldings
		asset.AveragePrice = asset.TotalValue / asset.TotalHolding
		portfolio.Assets[sym] = asset
		portfolio.TotalAssetsBalance = portfolio.TotalAssetsBalance + asset.TotalValue
	}

	portfolio.AssetCount = len(portfolio.Assets)
	return portfolio, nil
}
