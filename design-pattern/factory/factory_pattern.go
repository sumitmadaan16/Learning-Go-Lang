package factory

import "fmt"

type Vehicle interface {
	Drive()
}

// concrete types
type car struct{}

func (c car) Drive() {
	fmt.Println("Driving a car...")
}

type bike struct{}

func (b bike) Drive() {
	fmt.Println("Riding a bike...")
}

// factory impementation
func GetVehicle(vType string) Vehicle {
	switch vType {
	case "car":
		return car{}
	case "bike":
		return bike{}
	default:
		return nil
	}
}

func FactoryPattern() {
	car := GetVehicle("car")
	car.Drive()

	bike := GetVehicle("bike")
	bike.Drive()
}
