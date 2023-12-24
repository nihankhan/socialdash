// main.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nihankhan/socialdash/internal/api"
)

func main() {
	// Initialize and start Goroutine for auto-updating follower counts
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	//go startAutoUpdate(ctx)

	// Set up routes and start the server
	router := http.NewServeMux()
	router.HandleFunc("/", api.HomeHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Socialdash server is Running on 127.0.0.1:8080")

	done := make(chan struct{})

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}

		done <- struct{}{}
	}()

	<-done

}

// Graceful shutdown on interrupt signal
/*go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Println("Shutting down server gracefully...")
		cancel() // Signal the auto-update Goroutine to exit
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Error during server shutdown: %v\n", err)
		}
	}()

	// Start the server
	log.Println("Server listening on :8080...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %v\n", err)
	}
}


func startAutoUpdate(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Update follower counts
			updateFollowerCounts()
		case <-ctx.Done():
			// Context canceled, exit the Goroutine
			log.Println("Auto-update Goroutine exiting...")
			return
		}
	}
}

func updateFollowerCounts() {
	// Implement the logic to update follower counts for all platforms
	// You can call the respective services here
}
*/
