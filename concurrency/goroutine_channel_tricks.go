package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Main go routine starting")

	// we can use channels to make communications between the main go routine
	// and other go routines so than they can maintain synchronization
	ch := make(chan string, 1)
	go printRandomMessages3(ch)

	// Receive from ch, and assign value to a variable
	// The main go routine will be halted here until message received at the channel
	for i := 0; i < 10; i++ {
		// we can just receive the channel with <-ch and not assign the received value to
		// anywhere. It will serve for blocking purpose only.
		message := <-ch
		fmt.Print(message)
	}

	fmt.Println("Main go routine exiting")
}

func printRandomMessages3(ch chan string) {
	fmt.Println("printRandomMessages go routine starting")
	// It doesn't matter how many messages are sent via the channel ch
	// if the channel receives 10 messages in the main go routine
	// this for loop will not iterate more than 10 times
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("Message %v\n", i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}

	// this part of the code will remain unreachable
	// as soon as the 10th message will be sent via channel ch
	// main go routine will be executed without blocking and
	// thus the program will be finished.
	fmt.Println("printRandomMessages go routine exiting")
}
