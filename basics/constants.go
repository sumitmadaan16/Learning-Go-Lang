package main

import (
	"fmt"
)

func ConstantDemo() {
	fmt.Println()
	fmt.Println("Inside ConstantDemo")

	const a int = 12
	fmt.Println(a)
	var b float64 = float64(a)
	fmt.Println(b)

	//const expression
	const (
		length = 10
		width  = 20
		area   = length * width
	)
	fmt.Printf("area = %v , Type Area = %T\n", area, area) // %v gives values, and %T tells type , %t is for boolean

	// enumerated constant expresion
	const (
		red = iota
		blue
		green
	)
	fmt.Printf("Red: %v \nBlue: %v \nGreen: %v", red, blue, green)
}
