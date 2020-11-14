package main

import (
	"log"
	"net/http"

	"github.com/andresuchitra/inventory-backend/controllers"
	"github.com/andresuchitra/inventory-backend/models"
	"github.com/gorilla/mux"
)

var router *mux.Router

func initRouters() {
	router.HandleFunc("/cars", controllers.GetCars).Methods("GET")
	router.HandleFunc("/cars/{id}", controllers.GetCar).Methods("GET")
	router.HandleFunc("/cars", controllers.CreateCar).Methods("POST")
}

func main() {
	models.ConnectDB()

	router = mux.NewRouter().StrictSlash(true)
	initRouters()

	log.Fatal(http.ListenAndServe(":3001", router))
}
