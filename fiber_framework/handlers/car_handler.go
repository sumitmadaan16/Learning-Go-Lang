package handlers

import (
	"fiber-framework/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllCars(c *fiber.Ctx) error {
	allCars, err := models.GetAllCars()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal error"})
	}
	return c.JSON(allCars)
}

func GetById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Id"})
	}
	car, err := models.GetCarByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if car == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Car not found"})
	}
	return c.JSON(car)
}

func UpdateCar(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Id"})
	}
	var car models.Car
	if err := c.BodyParser(&car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})
	}
	car.Id = id
	if err := car.UpdateCar(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(car)
}

func CreateCars(c *fiber.Ctx) error {
	car := new(models.Car)
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON input",
		})
	}
	if err := car.Insert(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println("Car added to the inventory with ID:", car.Id)
	return c.Status(fiber.StatusCreated).JSON(car)
}

func DeleteCar(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Id"})
	}
	if err := models.DeleteCar(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
