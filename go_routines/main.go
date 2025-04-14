package main

import (
	"sync"
	"time"
)

func sayHello() {
	println("Hello, World!")
	time.Sleep(2 * time.Second) // Sleep for 2 seconds to simulate work
	println("Hello, World! after 2 seconds")
}

func syaHi(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	println("Hi, shafi!", i)
}

var wg sync.WaitGroup

func main() {
	println("Learning go routine and channels")
	for i := 1; i <= 5; i++ {
		wg.Add(1)        // Increment the WaitGroup counter
		go syaHi(i, &wg) // Start a goroutine for each iteration
	}
	wg.Wait() // Wait for all goroutines to finish
	println("All goroutines finished")
}
