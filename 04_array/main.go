package main

import "fmt"

func main() {
	fmt.Println("this file is for array")
	var fruits [2]string

	fruits[1] = "apple"
	fruits[0] = "banana"
	fmt.Println(fruits)

	var vegy = [3]int{4, 9, 10}
	fmt.Println(vegy)

}
