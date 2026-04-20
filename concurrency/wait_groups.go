package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup) { // this is a worker function that is being called from the demo func.
	defer wg.Done() // defer schedules the call to wg.Done() to run when the function returns, no matter how it exits (normal return, early return, or even after a panic if recovered).
	fmt.Printf(" worker %v started\n", i)
	time.Sleep(1 * time.Second) // adds a 1 sec. delay so that all the routines can be started
	fmt.Printf("worker %v completed\n", i)
}

func WaitGroupDemo() {
	var wg sync.WaitGroup // creating the instance of waitgroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i+1, &wg) // here we run 5 go routines parallely and these run sleep and gets completed
		//The Go runtime doesn’t enforce a hierarchy — goroutines are all peers. There’s no “tree” of goroutines like in some actor models.
	}
	wg.Wait()
	fmt.Println("Main worker completed")
}
