package main

import (
	// "fmt"
	"net/http"
	"urlShortener/src/logger"
	"urlShortener/src/router"
)

// main function is uded to call the API function and to start the http server for API
func main() {

	router.NewRouter()
	// fmt.Println("=================>")
	logger.Log.Println("Server started at 8082 port")
	logger.Log.Fatal(http.ListenAndServe(":8084", nil))
}
