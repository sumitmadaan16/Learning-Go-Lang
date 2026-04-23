package models

import (
	"cars/config"
	"database/sql"
	"fmt"
)

// Car struct with JSON tags
type Car struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Year  int     `json:"year"`
	Price float64 `json:"price"`
}

// Insert a new car
func (c *Car) Insert() error {
	query := `INSERT INTO cars(name, brand, year, price) 
              VALUES ($1, $2, $3, $4) RETURNING id`
	err := config.DB.QueryRow(query, c.Name, c.Brand, c.Year, c.Price).Scan(&c.Id)
	if err != nil {
		return fmt.Errorf("error inserting car: %w", err)
	}
	return nil
}

// Get all cars
func GetAllCars() ([]Car, error) {
	rows, err := config.DB.Query(`SELECT id, name, brand, year, price FROM cars`)
	if err != nil {
		return nil, fmt.Errorf("error fetching cars: %w", err)
	}
	defer rows.Close()

	var cars []Car
	for rows.Next() {
		var c Car
		if err := rows.Scan(&c.Id, &c.Name, &c.Brand, &c.Year, &c.Price); err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}

// Get car by ID
func GetCarByID(id int) (*Car, error) {
	row := config.DB.QueryRow(`SELECT id, name, brand, year, price FROM cars WHERE id=$1`, id)
	var c Car
	err := row.Scan(&c.Id, &c.Name, &c.Brand, &c.Year, &c.Price)
	if err == sql.ErrNoRows {
		return nil, nil // not found
	}
	if err != nil {
		return nil, fmt.Errorf("error fetching car: %w", err)
	}
	return &c, nil
}

// Update car by ID
func (c *Car) Update() error {
	query := `UPDATE cars SET name=$1, brand=$2, year=$3, price=$4 WHERE id=$5`
	_, err := config.DB.Exec(query, c.Name, c.Brand, c.Year, c.Price, c.Id)
	if err != nil {
		return fmt.Errorf("error updating car: %w", err)
	}
	return nil
}

// Delete car by ID
func DeleteCar(id int) error {
	query := `DELETE FROM cars WHERE id=$1`
	result, err := config.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting car: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking delete result: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no car found with id %d", id)
	}
	return nil
}
