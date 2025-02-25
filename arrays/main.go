package main

import "fmt"

func main() {
	fmt.Println("Wellcome to arrays")

	// myarray := [...]int{5, 4, 6, 8}

	var myarray [4]int
	myarray[0] = 2

	fmt.Println("the is the values of array", myarray)
}
