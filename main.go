package main

import (
	"Skool_Saver/src/config"
	"Skool_Saver/src/logger"
	router "Skool_Saver/src/routers"
	"net/http"
	"os"
)

func main() {
	// Initialize database connection
	db, err := config.ConnectMySqlDbSlaveSingleton()
	if err != nil {
		logger.Log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Set up the router
	router.NewRouter()

	// Get port from environment variable, fallback to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback to port 8080 if not set
	}

	// Start the server
	logger.Log.Printf("Server starting on port %s", port)
	logger.Log.Fatal(http.ListenAndServe(":"+port, nil))
}
