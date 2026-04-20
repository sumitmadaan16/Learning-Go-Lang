package main

import (
	"fmt"
	"time"
)

func GoRoutineDemo() {
	go counter()                 // when we simply write this statement then it does not wait for this function to execute it just executes the main stack and ends the program
	time.Sleep(12 * time.Second) // but when we add a 10 sec. delay it starts executing the other go routines until those 10 sec.
	fmt.Println("Inside the main go routine")
}

func counter() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second) // just to let it wait for a sec. to see it printing in console and also to see at what sec. does the main stack gets executed
		fmt.Println(i)
	}
	fmt.Println("Inside the seperate Go Routine")
}
