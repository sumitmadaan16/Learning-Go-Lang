package handlers

import (
	"cars/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func CarHandler(writeRes http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	entity := strings.TrimPrefix(path, "/cars")
	entity = strings.Trim(entity, "/")

	switch req.Method {
	case "GET":
		if entity == "" {
			// Fetch all cars from DB
			cars, err := models.GetAllCars()
			if err != nil {
				http.Error(writeRes, err.Error(), http.StatusInternalServerError)
				return
			}
			writeRes.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writeRes).Encode(cars)
		} else {
			id, err := strconv.Atoi(entity)
			if err != nil {
				http.Error(writeRes, "Invalid ID", http.StatusBadRequest)
				return
			}
			car, err := models.GetCarByID(id)
			if err != nil {
				http.Error(writeRes, err.Error(), http.StatusInternalServerError)
				return
			}
			if car == nil {
				http.Error(writeRes, "Car not found", http.StatusNotFound)
				return
			}
			writeRes.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writeRes).Encode(car)
		}

	case "PUT":
		if entity == "" {
			http.Error(writeRes, "Provide an ID to update", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(entity)
		if err != nil {
			http.Error(writeRes, "Invalid ID", http.StatusBadRequest)
			return
		}
		// Decode JSON body into Car struct
		car := &models.Car{Id: id}
		if err := json.NewDecoder(req.Body).Decode(car); err != nil {
			http.Error(writeRes, "Invalid JSON input", http.StatusBadRequest)
			return
		}
		// Call Update method
		if err := car.Update(); err != nil {
			http.Error(writeRes, err.Error(), http.StatusInternalServerError)
			return
		}
		writeRes.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writeRes).Encode(car)

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
			deleteCar(writeRes, req, id)
		}
	}

}

func createCars(writeRes http.ResponseWriter, req *http.Request) {
	car := &models.Car{}
	if err := json.NewDecoder(req.Body).Decode(car); err != nil {
		http.Error(writeRes, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	if err := car.Insert(); err != nil {
		http.Error(writeRes, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Car added to the inventory with ID: ", car.Id)
	writeRes.Header().Set("Content-Type", "application/json")
	writeRes.WriteHeader(http.StatusCreated)
	json.NewEncoder(writeRes).Encode(car)
}
func deleteCar(writeRes http.ResponseWriter, req *http.Request, id int) {
	err := models.DeleteCar(id)
	if err != nil {
		http.Error(writeRes, err.Error(), http.StatusNotFound)
		return
	}
	writeRes.WriteHeader(http.StatusNoContent) // 204 No Content
}
