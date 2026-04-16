package main

import (
	"fmt"
	"sync"
)

// global variables
var (
	basicCounter int
	mutex        sync.Mutex
)

// basic increment function
func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock() // mutex lock enabled
	basicCounter++
	mutex.Unlock() // resources freed
}

func BasicMutexDemo() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", basicCounter)
}
