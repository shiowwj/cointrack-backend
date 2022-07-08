package routes

import (
	"github.com/gorilla/mux"
	"github.com/shiowwj/go-cointracker-crud/pkg/controllers"
)

// TODO: add validators to check incoming response
var RegisterPortfolioAssetsRoutes = func(r *mux.Router) {

	// portfolio endpoint
	r.HandleFunc("/assets", controllers.GetAssets).Methods("GET").Queries("uuid", "{uuid}")

	// Txns endpoint
	// r.HandleFunc("/transactions/all", controllers.GetAllTransactions).Methods("GET")
	// r.HandleFunc("/transaction/{transactionId}", controllers.GetTransactionById).Methods("GET")
	// r.HandleFunc("/transactions", controllers.GetTransactionsBySymbol).Methods("GET").Queries("symbol", "{symbol}")
	r.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions/{transactionId}", controllers.UpdateTransaction).Methods("PUT")
	r.HandleFunc("/transactions/{transactionId}", controllers.DeleteTransaction).Methods("DELETE")
}
