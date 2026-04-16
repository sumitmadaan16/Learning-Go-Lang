package main

import (
	"fmt"
	"sync"
)

var (
	mcounter int
	rwMutex  sync.RWMutex // read write lock
	wg       sync.WaitGroup
)

func readCounter() {
	defer wg.Done()
	rwMutex.RLock()
	fmt.Printf("Reading Counter: %d\n", mcounter)
	rwMutex.RUnlock()
}

func writeCounter() {
	defer wg.Done()
	rwMutex.Lock()
	mcounter++
	rwMutex.Unlock()
}

func ReadWriteMutexDemo() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writeCounter()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readCounter()
	}

	wg.Wait()
}
