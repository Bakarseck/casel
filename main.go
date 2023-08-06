package main

import (
	handlers "casel/lib/handlers"
	"log"
	"net/http"
)

func main() {

	// Handle routes
	http.HandleFunc("/", handlers.IndexHandler)

	// Load css
	fsAssets := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsAssets))

	// Log
	log.Println("Start Server on Port 8080")
	log.Println("http://localhost:8080")

	// Start Listening and serving
	http.ListenAndServe(":8080", nil)
}
