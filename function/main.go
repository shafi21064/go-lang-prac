package main

import "fmt"

func main() {
	//fmt.Println(adder(10, 20))
	result, error := divider(20, 40)

	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Result:", result)
	}
}

func adder(a, b int) int {
	return a + b
}

func divider(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't divided by zero")
	}
	return a / b, nil
}
