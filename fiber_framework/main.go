package main

import (
	"fiber-framework/config"
	"fiber-framework/handlers"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

/* the basic of flow of a fiber application is

Client (HTTP request)
       ↓
Fiber Server (app.Listen)
       ↓
Router in main.go (method + path)
       ↓
Middleware (auth, logging, limiter, etc.)
       ↓
Handler (business logic)
       ↓
Model / Service (DB operations)
       ↓
Database
       ↓
Response (JSON / HTML / text)
       ↓
Client

*/

func main() {
	config.ConnectDB() // connection to DataBase

	app := fiber.New() // entry point to fiber framework. this basically creates *fiber.Ctx struct

	app.Use(logger.New()) // using fiber's built-in middleware's logger

	app.Use(basicauth.New(basicauth.Config{ // this is how u enable basic Auth in your application, first we call the basicauth's constructor then we set the configurations in configuration we have a map which we name as Users
		Users: map[string]string{
			"Admin": "12345",
		},
	}))

	// fiber directly handles the http methods
	app.Post("/cars", handlers.CreateCars)
	app.Put("/cars/:id", handlers.UpdateCar)
	app.Delete("/cars/:id", handlers.DeleteCar)
	app.Get("/cars", handlers.GetAllCars)
	app.Get("/cars/:id", handlers.GetById)

	fmt.Println("server started")
	err := app.Listen(":8080") // opening a port to host our website
	if err != nil {
		fmt.Println("server stopped abruptly..")
		return
	}
}
