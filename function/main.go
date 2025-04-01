package main

import "fmt"

func main() {
	//fmt.Println(adder(10, 20))
	// result, error := divider(20, 40)

	// if error != nil {
	// 	fmt.Println("Error:", error)
	// } else {
	// 	fmt.Println("Result:", result)
	// }

	fmt.Println(sum(1, 2, 3, 4, 5))
}

// Single return function
func adder(a, b int) int {
	return a + b
}

// Multi return function
func divider(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't divided by zero")
	}
	return a / b, nil
}

// Named return function
func add(a, b int) (sum int) {
	sum = a + b
	return
}

// Variadic function
func sum(num ...int) int {
	total := 0
	for _, item := range num {
		total += item
	}
	return total
}
