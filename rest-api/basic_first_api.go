package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Car struct {
	Id    int64
	Name  string
	Brand string
	Year  int
	Price float64
}

var Cars = make(map[int64]Car) // this is a map that acts as a database

var mu sync.Mutex //

//HTTP methods

// GET	/cars
// POST /cars
// DELETE	/cars/:id

func carHandler(writeRes http.ResponseWriter, req *http.Request) {
	path := req.URL.Path // fetches path from the URL

	entity := strings.TrimPrefix(path, "/cars")
	entity = strings.Trim(entity, "/")

	//r.URL.Path gives the request path (e.g., /cars/123).
	//TrimPrefix removes /cars, leaving just the ID part.
	//Trim cleans up slashes so you get "123" instead of "/123/".

	switch req.Method { // creating a switch case for intercepting the requests
	case "GET":
		if entity == "" {
			var allCars []Car
			fmt.Println("Fetching all cars...")
			for _, car := range Cars {
				allCars = append(allCars, car)
			}
			json.NewEncoder(writeRes).Encode(allCars)
		} else {
			id, err := strconv.Atoi(entity)
			if err != nil {
				http.Error(writeRes, "Invalid ID", http.StatusBadRequest)
				return
			}
			getCar(writeRes, id)
		}

	case "POST":
		if entity == "" {
			createCars(writeRes, req)
		} else {
			http.Error(writeRes, "Incorrect post request", http.StatusBadRequest)
		}

	case "DELETE":
		if entity == "" {
			http.Error(writeRes, "API not supported. Enter a proper ID", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)
			deleteCar(writeRes, id)
		}
	}

}

func createCars(writeRes http.ResponseWriter, req *http.Request) { //method to Add a new car with an id
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Car added...")
	var car Car                                                    //creating a car obj
	if err := json.NewDecoder(req.Body).Decode(&car); err != nil { //fetching the car object from the request and parsing it into json format
		http.Error(writeRes, "Invalid JSON input", http.StatusBadRequest) // if there is an error we return invalid json input in response
		return
	}

	id := rand.Intn(5000)                                     // randomly generating an id between 1 and 5000
	car.Id = int64(id)                                        // set id to car
	Cars[car.Id] = car                                        // fetch the id and update the data in map/database
	writeRes.Header().Set("Content-Type", "application/json") //setting the content type so that data flow remains consistent
	writeRes.WriteHeader(http.StatusCreated)                  // add response to it
	json.NewEncoder(writeRes).Encode(car)
}

func getCar(writeRes http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Fetching car with Id...")
	val, ok := Cars[int64(id)]
	if !ok {
		http.Error(writeRes, "Car not found", http.StatusBadRequest)
		return
	}

	writeRes.Header().Set("Content-Type", "application/json")
	writeRes.WriteHeader(http.StatusOK)
	json.NewEncoder(writeRes).Encode(val)
}

func deleteCar(writeRes http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Deleting a car")
	_, ok := Cars[int64(id)]
	if !ok {
		http.Error(writeRes, "Car not found", http.StatusBadRequest)
		return
	}
	delete(Cars, int64(id))
	writeRes.WriteHeader(http.StatusNoContent)
}
