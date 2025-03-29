package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// myMap := map[string]int{
	// 	"apple":  10,
	// 	"banana": 20,
	// }

	myLanguages := make(map[string]any)

	myLanguages["js"] = "javascript"
	myLanguages["score"] = 100

	fmt.Println(myLanguages)
}
