package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Main go routine starting")

	// this goroutine might be called but will not be executed
	// the main go routine will not wait for this go routine to be
	// executed and finished by itself
	// as main go routine and printRandomMessages go routine runs separately
	go printRandomMessages1()

	// We can add some blocking conditions for main go routine like -
	// for {}
	// But this will block main go routine from executing also
	// Or we can use channels go make communications between the main go routine
	// and other go routines so than they can maintain synchronization

	fmt.Println("Main go routine exiting")
}

func printRandomMessages1() {
	fmt.Println("printRandomMessages go routine starting")
	for i := 0; i < 10; i++ {
		fmt.Printf("Message %v\n", i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
	fmt.Println("printRandomMessages go routine exiting")
}
