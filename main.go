package main

import (
	"Skool_Saver/src/logger"
	router "Skool_Saver/src/routers"
	"net/http"
	"os"
)

func main() {
	router.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback to 8080 if PORT is not set
	}

	logger.Log.Printf("Server starting on port %s", port)
	logger.Log.Fatal(http.ListenAndServe(":"+port, nil))
}