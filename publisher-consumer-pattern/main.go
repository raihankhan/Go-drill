package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func keepProcessAlive() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP)
	<-quit
	fmt.Println("Terminating main thread")
}

func main() {
	var publisherWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	buffer := make(chan int, 5)

	publisherFunc := func(wg *sync.WaitGroup, buffer chan int) {
		defer wg.Done()
		for data := 1; data <= 10; data++ {
			buffer <- data
			fmt.Printf("published data : %d\n", data)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		}
		fmt.Println("Producer session Terminates")
	}

	consumerFunc := func(wg *sync.WaitGroup, buffer chan int) {
		defer wg.Done()
		for data := range buffer {
			fmt.Printf("Consumed data : %d\n", data)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		}
		fmt.Println("Consumer session Terminates")
	}

	terminateChannel := func(wg *sync.WaitGroup, buffer chan int) {
		wg.Wait()
		close(buffer)
		fmt.Println("Closed channel")
	}

	publisherWg.Add(1)
	go publisherFunc(&publisherWg, buffer)

	consumerWg.Add(1)
	go consumerFunc(&consumerWg, buffer)

	go terminateChannel(&publisherWg, buffer)

	keepProcessAlive()
}
