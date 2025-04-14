package main

import (
	"math/rand"
	"sync"
	"time"
)

func generateNumber() int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(100)
}

func syaHi(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch { // Keep reading until channel is closed
		println("Hi, shafi!", num)
	}
}

func channelNumber(ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	defer close(ch) // Close channel when done sending
	for i := 1; i <= 5; i++ {
		ch <- i // Send a random number to the channel
	}
}

func main() {
	println("learn go channels")

	var wg sync.WaitGroup

	chanNumber := make(chan int)

	wg.Add(2) // Increment the WaitGroup counter

	go syaHi(chanNumber, &wg)         // Start a goroutine for each iteration
	go channelNumber(chanNumber, &wg) // Start a goroutine for each iteration

	wg.Wait() // Wait for all goroutines to finis

}
