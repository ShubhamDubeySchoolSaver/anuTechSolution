package main

import (
	"Skool_Saver/src/logger"
	router "Skool_Saver/src/routers"
	"net/http"
)

func main() {
	router.NewRouter()
	logger.Log.Println("Server started at 8080 port")
	logger.Log.Fatal(http.ListenAndServe(":8080", nil))
}
