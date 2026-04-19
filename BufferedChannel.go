package main

import (
	"fmt"
	"time"
)

func BufferedChannelDemo() {
	ch := make(chan int, 2) // channel created with a buffer size of 2

	go func() {
		ch <- 1
		fmt.Println("Message 1 sent")
		ch <- 2
		fmt.Println("Message 2 sent")
		ch <- 3
		fmt.Println("Message 3 sent")
	}()

	fmt.Println("Waiting for sender...")
	time.Sleep(2 * time.Second)

	msg1 := <-ch
	fmt.Println("Message received", msg1)
	msg2 := <-ch
	fmt.Println("Message received", msg2)
	msg3 := <-ch
	fmt.Println("Message received", msg3)
	close(ch)
}
