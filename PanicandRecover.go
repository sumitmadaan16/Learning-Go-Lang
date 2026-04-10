package main

import "fmt"

func divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered : ", r)
		}
	}()
	if b == 0 {
		panic("Divided By zero")
	}
	return a / b
}

func Panicdemo() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered : ", r)
	//	}
	//}()

	//if we keep this defer block here then it will not run the line of code after where the panic happened,
	//but if we keep it above it will run after that also
	fmt.Println(divide(4, 2))
	fmt.Println(divide(4, 0)) // Panicked Here therefore Not preceding ahead
	//prints zero cuz the default value for int, since the return expression was interrupted by panic).

	fmt.Println("program continues running") // execution coninues here cuz the panic was handled inside the divide func itself,
	// and the execution continues in the caller after the function call.
	fmt.Println(divide(17, 2))
}

// Stack Unwinding
func A() {
	defer fmt.Println("A defer")
	B()
}

func B() {
	defer fmt.Println("B defer")
	C()
}

func C() {
	defer fmt.Println("C defer")
	panic("Boom!")
}

func StackUnwindingDemo() {
	A()
}
