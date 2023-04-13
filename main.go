package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"github.com/9ziggy9/go-starter/config"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	// LOAD ENV AND SET PORT
	if err := config.LoadEnv(".env"); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file: %s\n", err)
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Fprintf(os.Stderr, "Error loading port env variable.\n")
		os.Exit(1)
	}

	// START ROUTER
	r := mux.NewRouter()

	// ENDPOINTS
	r.HandleFunc("/api/message", messageHandler).Methods(http.MethodGet)

	// LISTEN AND SERVE
	fmt.Println("Opening port at "+port+" and listening...")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Text: "Hello, World!",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %s\n", err)
	}
}
