package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in go")

	rand.Seed(time.Now().Unix())
	number := rand.Intn(6) + 1
	fmt.Println("Generated number", number)

	switch number {
	case 1:
		fmt.Println("This is 1")
	case 2:
		fmt.Println("This is 2")
	case 3:
		fmt.Println("This is 3")
	case 4:
		fmt.Println("This is 4")
	case 5:
		fmt.Println("This is 5")
	case 6:
		fmt.Println("This is 6")

	default:
		fmt.Println("What is this!")
	}
}
