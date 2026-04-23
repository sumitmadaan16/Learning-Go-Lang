package main

import (
	"cars/config"
	"cars/handlers"
	"fmt"
	"net/http"
)

func main() {
	config.ConnectDB()
	http.HandleFunc("/cars", handlers.CarHandler)
	http.HandleFunc("/cars/", handlers.CarHandler)
	fmt.Println("server started")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server stopped..")
		return
	}
}
