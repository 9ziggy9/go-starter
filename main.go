package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"github.com/9ziggy9/go-starter/config"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
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

	// OPEN DB CONNECTION
	dsn := config.BuildDSN(
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PORT"),
	)

	fmt.Printf("\n\n\n CONNECTING TO DB VIA DSN\n %s\n\n\n", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening DB connection: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(db)
	
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
