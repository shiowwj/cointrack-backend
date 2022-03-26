package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/shiowwj/go-cointracker-crud/pkg/models"
	"github.com/shiowwj/go-cointracker-crud/pkg/utils"
)

var NewHolding models.Holding

func GetHoldingsAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s\n", r.Method, r.RequestURI)
	newHoldings := models.GetAllHoldings()
	res, _ := json.Marshal(newHoldings)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetHoldingById(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s\n", r.Method, r.RequestURI)
	vars := mux.Vars(r)
	holdingId := vars["holdingId"]
	ID, err := strconv.ParseInt(holdingId, 0, 0)
	if err != nil {
		log.Println("Error while parsing")
	}
	holdingDetails, _ := models.GetHoldingById(ID)
	res, _ := json.Marshal(holdingDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetHoldingsBySymbol(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s\n", r.Method, r.RequestURI)
	vars := mux.Vars(r)
	symbol := strings.ToUpper(vars["symbol"])
	holdingDetails, _ := models.GetHoldingsBySymbol(symbol)
	res, _ := json.Marshal(holdingDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateHolding(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s\n", r.Method, r.RequestURI)
	CreateHolding := &models.Holding{}
	utils.ParseBody(r, CreateHolding)
	h := CreateHolding.CreateHolding()
	res, _ := json.Marshal(h)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteHolding(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s\n", r.Method, r.RequestURI)
	vars := mux.Vars(r)
	holdingId := vars["holdingId"]
	ID, err := strconv.ParseInt(holdingId, 0, 0)
	if err != nil {
		log.Println("Error while parsing")
	}
	holding := models.DeleteHolding(ID)
	res, _ := json.Marshal(holding)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateHolding(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s\n", r.Method, r.RequestURI)
	var updateHolding = &models.Holding{}
	utils.ParseBody(r, updateHolding)
	vars := mux.Vars(r)
	holdingId := vars["holdingId"]
	ID, err := strconv.ParseInt(holdingId, 0, 0)
	if err != nil {
		log.Println("Error while parsing")
	}
	holdingDetails, db := models.GetHoldingById(ID)

	if updateHolding.Symbol != "" {
		holdingDetails.Symbol = updateHolding.Symbol
	}
	if updateHolding.AmountBought != 0 {
		holdingDetails.AmountBought = updateHolding.AmountBought
	}
	if updateHolding.Currency != "" {
		holdingDetails.Currency = updateHolding.Currency
	}
	if updateHolding.BoughtAtMarketPrice != 0 {
		holdingDetails.BoughtAtMarketPrice = updateHolding.BoughtAtMarketPrice
	}
	if updateHolding.BoughtAtDate != "" {
		holdingDetails.BoughtAtDate = updateHolding.BoughtAtDate
	}
	if updateHolding.SoldAtMarketPrice != 0 {
		holdingDetails.SoldAtMarketPrice = updateHolding.SoldAtMarketPrice
	}
	if updateHolding.SoldAtDate != "" {
		holdingDetails.SoldAtDate = updateHolding.SoldAtDate
	}
	if updateHolding.Status != "" {
		holdingDetails.Status = updateHolding.Status
	}

	db.Save(&holdingDetails)
	res, _ := json.Marshal(holdingDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
