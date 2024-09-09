package main

import (
	"fmt"
	"time"
)

func main() {
	stringChannel := make(chan string)
	intChannel := make(chan int)

	// Goroutine to send a string message after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		stringChannel <- "Hello, world!"
	}()

	// Goroutine to send an integer after 1 second
	go func() {
		time.Sleep(2 * time.Second)
		intChannel <- 4245
	}()

	// Using select{} to listen to the channels
	select {
	case msg := <-stringChannel:
		fmt.Println("Received string message:", msg)
	case num := <-intChannel:
		fmt.Println("Received integer:", num)
	case <-time.After(3 * time.Second): // Timeout after 3 seconds
		fmt.Println("No messages received in time.")
	}
}
