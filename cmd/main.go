package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nihankhan/socialdash/internal/api"
)

func main() {
	// Set up routes and start the server
	router := http.NewServeMux()
	router.HandleFunc("/", api.HomeHandler)

	// Serve static files
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Socialdash server is Running on 127.0.0.1:8080")

	// Graceful shutdown on interrupt signal
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
