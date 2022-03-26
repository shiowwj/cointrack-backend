package routes

import (
	"github.com/gorilla/mux"
	"github.com/shiowwj/go-cointracker-crud/pkg/controllers"
)

var RegisterPortfolioHoldingsRoutes = func(r *mux.Router) {
	r.HandleFunc("/holdings/all", controllers.GetHoldingsAll).Methods("GET")
	r.HandleFunc("/holding/{holdingId}", controllers.GetHoldingById).Methods("GET")
	r.HandleFunc("/holdings", controllers.GetHoldingsBySymbol).Methods("GET").Queries("symbol", "{symbol}")
	r.HandleFunc("/holdings", controllers.CreateHolding).Methods("POST")
	r.HandleFunc("/holdings/{holdingId}", controllers.UpdateHolding).Methods("PUT")
	r.HandleFunc("/holdings/{holdingId}", controllers.DeleteHolding).Methods("DELETE")
}
