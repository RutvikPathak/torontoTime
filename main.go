package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

func getTime(w http.ResponseWriter, r *http.Request) {
	// Set the timezone to toronto

	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Get the current time in toronto timezone
	currentTime := time.Now().In(loc)

	// Create the response struct
	response := TimeResponse{
		CurrentTime: currentTime.Format("2006-1-2 15:4:5"),
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the struct to JSON and write the response
	json.NewEncoder(w).Encode(response)

}

func main() {
	http.HandleFunc("/time", getTime)
	fmt.Println("Server is running on server port :3000")
	http.ListenAndServe(":3000", nil)

}
