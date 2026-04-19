package main

import (
	"fmt"
	"sync"
	"time"
)

func workerPool(i int, Tasks <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range Tasks {
		fmt.Printf("Task %d picked up by worker %d \n", task, i)
		time.Sleep(1 * time.Second)
		result <- task * task
	}
}

func ConcurrencyDemo() {
	const workers = 3
	const jobs = 8

	taskBuffer := jobs - workers

	tasks := make(chan int, taskBuffer)
	results := make(chan int, jobs)

	var wg sync.WaitGroup

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go workerPool(i, tasks, results, &wg)
	}
	for i := 1; i <= jobs; i++ {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println(r)
	}
}
