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
	ch := make(chan bool)
	go printRandomMessages2(ch)

	// Receive from ch, and assign value to a variable
	// The main go routine will be halted here until message received at the channel
	isRandomMessagesPrinted := <-ch
	if isRandomMessagesPrinted {
		fmt.Println("printRandomMessages2 go routine finished executing successfully")
	}

	// we can just receive the channel with <-ch and not assign the received value to
	// anywhere. It will serve for blocker only purpose

	fmt.Println("Main go routine exiting")
}

func printRandomMessages2(ch chan bool) {
	fmt.Println("printRandomMessages go routine starting")
	for i := 0; i < 10; i++ {
		fmt.Printf("Message %v\n", i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
	fmt.Println("printRandomMessages go routine exiting")
	ch <- true // Send bool value to channel ch as our go routine is finished executing
}
