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
	router.HandleFunc("/cars/{id}", controllers.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars", controllers.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", controllers.DeleteCar).Methods("DELETE")
}

func main() {
	models.ConnectDB()

	router = mux.NewRouter().StrictSlash(true)
	initRouters()

	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, router)

	// setup handlers with CORS
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServeTLS(":3001", os.Getenv("CRT_FILE"), os.Getenv("KEY_FILE"), handlers.CORS(originsOK, methodsOK)(loggedRouter)))
}
