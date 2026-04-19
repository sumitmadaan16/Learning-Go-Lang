package main

import (
	"fmt"
	"time"
)

func send(ch chan<- int) { // sending channel
	fmt.Println("Sending message onto channel")
	time.Sleep(1 * time.Second)
	ch <- 16
}
func receive(ch <-chan int) { // this is a receiving channel
	time.Sleep(2 * time.Second) // added a sleep cuz this keeps the main go routine running for a bit longer until sender sends the message
	fmt.Println("waiting for the message")
	msg := <-ch
	fmt.Println("message received is: ", msg)
}
func ChannelDirectionDemo() {
	ch := make(chan int) // created a bidirectional channel for both sending and receiving

	go send(ch) // made this as a go routine as we want to send message from another go routine
	receive(ch) // why haven't we created this as a go routine?
	// cuz this has to be in main stack else the main stack would just gets executed and all the remaining go routines are deleted.
	close(ch)
}
