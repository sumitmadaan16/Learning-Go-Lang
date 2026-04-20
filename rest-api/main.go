package main

import (
	"fmt"
	"net/http"
)

func main() {
	//creating handelers to accept the incoming request
	http.HandleFunc("/cars", carHandler)
	http.HandleFunc("/cars/", carHandler)

	fmt.Println("server started")

	// opening a port to run server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server stopped..")
		return
	}

}
