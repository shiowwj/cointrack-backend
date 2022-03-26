package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shiowwj/go-cointracker-crud/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterPortfolioHoldingsRoutes(r)
	http.Handle("/", r)
	log.Println("Starting server at port...9010")
	err := http.ListenAndServe("localhost:9010", r)
	if err != nil {
		log.Fatal(http.ListenAndServe("localhost:9010", r))
	}
}
