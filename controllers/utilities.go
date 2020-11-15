package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type mapper map[string]interface{}

func respondJSON(w http.ResponseWriter, httpStatus int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(mapper{"data": payload})
}

func respondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// for internal server error, we hide the actual error message
	if code >= 500 {
		message = "Unexpected error ocurred. Please try again"
	} else {
		log.Println("REQUEST_ERROR -> " + message)
	}

	json.NewEncoder(w).Encode(mapper{"errors": message})
}

// validate param
func validateParam(vars map[string]string, w http.ResponseWriter) int {
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid Param: "+err.Error())
		return -1
	}

	return id
}
