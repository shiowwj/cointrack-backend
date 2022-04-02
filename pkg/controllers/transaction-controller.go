package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/shiowwj/go-cointracker-crud/pkg/models"
	"github.com/shiowwj/go-cointracker-crud/pkg/utils"
	"github.com/shiowwj/go-cointracker-crud/pkg/utils/log"
	"go.uber.org/zap"
)

var NewTransaction models.Transaction

// type Holding struct {
// }
func GetAssets(w http.ResponseWriter, r *http.Request) {
	log.Info("GetAssets", zap.Any("Method", r.Method), zap.Any("RequestURI", r.RequestURI))
	vars := mux.Vars(r)
	userId := strings.ToUpper(vars["uuid"])
	// TODO: Should have user authentication

	userAssets, err := models.GetUserAssets(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	res, _ := json.Marshal(userAssets)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	log.Info("GetAllTransactions", zap.Any("Method", r.Method), zap.Any("RequestURI", r.RequestURI))
	allTransactions := models.GetAllTransactions()
	res, _ := json.Marshal(allTransactions)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTransactionById(w http.ResponseWriter, r *http.Request) {
	log.Info("GetTransactionById", zap.Any("Method", r.Method), zap.Any("RequestURI", r.RequestURI))
	vars := mux.Vars(r)
	transactionId := vars["transactionId"]
	ID, err := strconv.ParseInt(transactionId, 0, 0)
	if err != nil {
		log.Info("Error while parsing")
	}
	txnDetails, _ := models.GetTransactionById(ID)
	res, _ := json.Marshal(txnDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func GetTransactionsBySymbol(w http.ResponseWriter, r *http.Request) {
// 	log.Info("GetTransactionsBySymbol", zap.Any("Method", r.Method), zap.Any("RequestURI", r.RequestURI))
// 	vars := mux.Vars(r)
// 	symbol := strings.ToUpper(vars["symbol"])
// 	txnDetails, _ := models.GetTransactionsBySymbol(symbol)
// 	res, _ := json.Marshal(txnDetails)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	log.Info("CreateTransaction", zap.Any("Method", r.Method), zap.Any("RequestURI", r.RequestURI))
	// body := json.Unmarshal(r.Body.Read(),)
	// log.Debug("CreateTransaction", zap.Any("r.Header", r.Header.Get("Uuid")))
	// get the
	CreateTxn := &models.Transaction{}
	utils.ParseBody(r, CreateTxn)
	t := CreateTxn.CreateTransaction(r.Header.Get("Uuid"))

	//TODO: Add failed creation http response
	res, _ := json.Marshal(t)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	log.Info("DeleteTransaction", zap.Any("Method", r.Method), zap.Any("RequestURI", r.RequestURI))
	vars := mux.Vars(r)
	transactionId := vars["transactionId"]
	ID, err := strconv.ParseInt(transactionId, 0, 0)
	if err != nil {
		log.Info("Error while parsing")
	}
	txn := models.DeleteTransaction(ID)

	res, _ := json.Marshal(txn)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	log.Info("UpdateTransaction", zap.Any("Method", r.Method), zap.Any("RequestURI", r.RequestURI))
	var updateTxn = &models.Transaction{}
	utils.ParseBody(r, updateTxn)
	vars := mux.Vars(r)
	transactionId := vars["transactionId"]
	ID, err := strconv.ParseInt(transactionId, 0, 0)
	if err != nil {
		log.Info("Error while parsing")
	}

	txnDetails, db := models.GetTransactionById(ID)
	if updateTxn.Type != "" {
		txnDetails.Type = updateTxn.Type
	}
	if updateTxn.Symbol != "" {
		txnDetails.Symbol = updateTxn.Symbol
	}
	if updateTxn.Name != "" {
		txnDetails.Name = updateTxn.Name
	}
	if updateTxn.Quantity != 0 {
		txnDetails.Quantity = updateTxn.Quantity
	}
	if updateTxn.Total != 0 {
		txnDetails.Total = updateTxn.Total
	}
	if updateTxn.PricePerCoin != 0 {
		txnDetails.PricePerCoin = updateTxn.PricePerCoin
	}
	if !updateTxn.DateTimeTxn.IsZero() {
		txnDetails.DateTimeTxn = updateTxn.DateTimeTxn
	}
	if updateTxn.Fees != 0 {
		txnDetails.Fees = updateTxn.Fees
	}

	db.Save(txnDetails)
	res, _ := json.Marshal(txnDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
