package helper

import (
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		log.Printf("Panic error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
