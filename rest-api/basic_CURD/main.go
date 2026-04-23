package main

import (
	"cars/config"
	"cars/handlers"
	"cars/middleware"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	config.ConnectDB()
	mux.HandleFunc("/cars", handlers.CarHandler)
	mux.HandleFunc("/cars/", handlers.CarHandler)

	loggedMux := middleware.Logger(mux)
	fmt.Println("server started")
	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		fmt.Println("server stopped..")
		return
	}
}
