package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Welcome to slices")

	fruitsSlices := []string{"mango", "banana", "avocado", "dates"}

	fmt.Println(fruitsSlices)

	fruitsSlices = append(fruitsSlices, "nuts", "papaya")
	fmt.Println(fruitsSlices)
	index := 2
	fruitsSlices = slices.Delete(fruitsSlices, index, index+1)
	fmt.Println(fruitsSlices)

}
