package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andresuchitra/inventory-backend/controllers"
	"github.com/andresuchitra/inventory-backend/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var router *mux.Router

func initRouters() {
	router.HandleFunc("/cars", controllers.GetCars).Methods("GET")
	router.HandleFunc("/cars/{id}", controllers.GetCar).Methods("GET")
	router.HandleFunc("/cars/{id}", controllers.UpdateCar).Methods("PUT", "OPTIONS")
	router.HandleFunc("/cars", controllers.CreateCar).Methods("POST", "OPTIONS")
	router.HandleFunc("/cars/{id}", controllers.DeleteCar).Methods("DELETE", "OPTIONS")
}

func main() {
	models.ConnectDB()

	router = mux.NewRouter().StrictSlash(true)
	initRouters()

	// setup handlers with CORS
	cors := handlers.CORS(
    handlers.AllowedHeaders([]string{"content-type"}),
    handlers.AllowedOrigins([]string{"*"}),
    handlers.AllowedMethods([]string{"OPTIONS", "GET", "POST", "PUT", "DELETE"}),
    handlers.AllowCredentials(),
	)

	router.Use(cors)

	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":3001", loggedRouter))
}
