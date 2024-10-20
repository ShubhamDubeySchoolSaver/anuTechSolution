package main

import (
    "Skool_Saver/src/config"
    "Skool_Saver/src/logger"
    router "Skool_Saver/src/routers"
    "net/http"
    "os"
)

func main() {
    // Log the deployment starting
    logger.Log.Println("Starting deployment...")

    // Initialize the database connection (and execute SQL script)
    db, err := config.ConnectMySqlDbSlaveSingleton()
    if err != nil {
        // Log if there is an issue with the database connection
        logger.Log.Fatalf("Error connecting to the database: %v", err)
    }
    defer db.Close()

    // Log successful database connection
    logger.Log.Println("Database connection established")

    // Initialize the router
    router.NewRouter()

    // Set up the server port
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Fallback to port 8080 if not set
    }

    // Log the server startup
    logger.Log.Printf("Server starting on port %s", port)

    // Start the HTTP server and log any fatal errors
    logger.Log.Fatal(http.ListenAndServe(":"+port, nil))
}
