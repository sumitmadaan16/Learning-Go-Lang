package main

import (
	"fmt"
	"sync"
	"time"
)

func Timer(ticker *time.Ticker, HeartbeatCount int, response chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= HeartbeatCount; i++ {
		<-ticker.C
		response <- "Dhak Dhak.."
	}
	close(response)
	ticker.Stop()
}

func TimeTickerDemo() {
	const HeartbeatCount = 50
	res := make(chan string, HeartbeatCount)
	ticker := time.NewTicker(100 * time.Millisecond)

	var wg sync.WaitGroup

	wg.Add(1)
	go Timer(ticker, HeartbeatCount, res, &wg)

	for response := range res {
		fmt.Println(response)
	}
	wg.Wait()
}
