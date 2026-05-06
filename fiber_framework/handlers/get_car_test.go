package handlers

import (
	"encoding/json"
	"fiber-framework/config"
	"fiber-framework/models"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func setup() *fiber.App {
	app := fiber.New()
	//app.Post("/cars", CreateCars)
	//app.Put("/cars/:id", UpdateCar)
	//app.Delete("/cars/:id", DeleteCar)  can be added if you test for all these methods
	app.Get("/cars", GetAllCars)
	app.Get("/cars/:id", GetById)
	return app
}

func TestMain(m *testing.M) {
	config.ConnectDB()
	m.Run()
}

// function to add the sample car just for test purpose
func sampleCar(t *testing.T) models.Car {
	t.Helper()
	car := models.Car{
		Name:  "Test Car",
		Brand: "Test Brand",
		Year:  2024,
		Price: 25000.00,
	}
	if err := car.Insert(); err != nil {
		t.Fatalf("seedCar: failed to insert test car: %v", err)
	}
	return car
}

// function to remove the sample test car from DB
func cleanupCar(t *testing.T, id int) {
	t.Helper()
	if err := models.DeleteCar(id); err != nil {
		t.Logf("cleanupCar: warning — could not delete car %d: %v", id, err)
	}
}

// this is a sample test block to get the cars by id
func TestGetById(t *testing.T) {
	app := setup()
	t.Run("Returns 200 and correct car with valid id", func(t *testing.T) {
		car := sampleCar(t)
		defer cleanupCar(t, car.Id)
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cars/%d", car.Id), nil)
		res, err := app.Test(req)
		assertNoErr(t, err)
		assertStatus(t, res, http.StatusOK)
		var received models.Car
		decodeJSON(t, res.Body, &received)

		if received.Id != car.Id {
			t.Errorf("received car id : %d, and expected id is: %d", received.Id, car.Id)
		}
	})

	t.Run("Returns 404 fro a non existant id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/cars/999999999", nil)
		res, err := app.Test(req)
		assertNoErr(t, err)
		assertStatus(t, res, http.StatusNotFound)
	})

	t.Run("return 400 for a non neumeric id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/cars/abcd", nil)
		res, err := app.Test(req)
		assertNoErr(t, err)
		assertStatus(t, res, http.StatusBadRequest)
	})

}

// getting all cars
func TestGetAllCars(t *testing.T) {
	app := setup()
	t.Run("Retrun 200 wiht a non-empty list of cars", func(t *testing.T) {
		car := sampleCar(t)
		defer cleanupCar(t, car.Id)
		req := httptest.NewRequest(http.MethodGet, "/cars", nil)
		res, err := app.Test(req)
		assertNoErr(t, err)
		assertStatus(t, res, http.StatusOK)

		var cars []models.Car
		decodeJSON(t, res.Body, &cars)

		if len(cars) == 0 {
			t.Error("Ex[ected atleast one car after adding a sample car")
		}
	})
}

// similarly you can do for adding , updating or deleing a car

// assertions in functions to reduce the boilerplate code
func assertNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertStatus(t *testing.T, resp *http.Response, want int) {
	t.Helper()
	if resp.StatusCode != want {
		t.Errorf("expected status %d, got %d", want, resp.StatusCode)
	}
}

func decodeJSON(t *testing.T, body io.Reader, v any) {
	t.Helper()
	if err := json.NewDecoder(body).Decode(v); err != nil {
		t.Fatalf("failed to decode response JSON: %v", err)
	}
}
