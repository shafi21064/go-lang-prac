package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	result := add(&a, &b)
	fmt.Println(result)
	//fmt.Println(&x)
}

func add(a, b *int) int {
	*a = 20
	return *a + *b
}
