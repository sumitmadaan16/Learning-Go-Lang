package handlers

import (
	"fiber-framework/config"
	"fiber-framework/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/gofiber/fiber/v2"
)

func BenchmarkGetById(b *testing.B) {
	config.ConnectDB()
	app := setup()

	// Seed once before the benchmark loop — we don't want DB inserts
	// skewing the handler timing

	car := models.Car{Name: "Bench Car", Brand: "Bench Brand", Year: 2024, Price: 10000}
	if err := car.Insert(); err != nil {
		b.Fatalf("failed to seed car: %v", err)
	}
	b.Cleanup(func() {
		err := models.DeleteCar(car.Id)
		if err != nil {
			return
		}
	})

	url := fmt.Sprintf("/cars/%d", car.Id)
	b.ResetTimer() // start timing only from here — excludes setup above

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, url, nil)
		resp, err := app.Test(req)
		if err != nil {
			b.Fatalf("request failed: %v", err)
		}
		err = resp.Body.Close()
		if err != nil {
			return
		} // always drain and close or you'll leak connections
	}
}
