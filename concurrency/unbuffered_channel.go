package main

import (
	"fmt"
	"time"
)

func UnbufferedChannelDemo() {
	ch := make(chan string) // channel created

	go func() { // another go routine invoked through this anonymous function call
		time.Sleep(2 * time.Second) // put sleep for few seconds to wait for message
		ch <- "Hello Learner"       // push the message onto channel
	}()
	fmt.Println("Waiting for message...") // back in main function stack and this line is executed automatically without the wait of another go routine to execute
	msg := <-ch                           // message received after 2 seconds of delay due to sleep
	fmt.Println(msg)                      // message printed

	close(ch)
}
