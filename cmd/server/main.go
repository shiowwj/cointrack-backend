package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/shiowwj/go-cointracker-crud/pkg/middleware"
	"github.com/shiowwj/go-cointracker-crud/pkg/routes"
	"github.com/shiowwj/go-cointracker-crud/pkg/utils/log"
	"go.uber.org/zap"
)

func main() {

	r := mux.NewRouter()
	r.Use(middleware.Middleware)
	routes.RegisterPortfolioAssetsRoutes(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
	})
	handler := c.Handler(r)
	// http.Handle("/", handler)
	log.Info("Starting server at port...9010")
	err := http.ListenAndServe("localhost:9010", handler)
	if err != nil {
		log.Fatal("Failed to start service", zap.Error(err))
	}
}

// router := mux.NewRouter()
// router.HandleFunc("/signup", ac.SignUp).Methods("POST")
// router.HandleFunc("/signin", ac.SignIn).Methods("POST")

// c := cors.New(cors.Options{
//     AllowedOrigins: []string{"http://localhost:8000"},
//     AllowCredentials: true,
// })

// handler := c.Handler(router)
// log.Fatal(http.ListenAndServe(":3000", handler)
