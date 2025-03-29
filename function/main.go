package main

import "fmt"

func main() {
	fmt.Println(adder(10, 20))
}

func adder(a, b int) int {
	return a + b
}
