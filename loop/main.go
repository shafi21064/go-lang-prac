package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednessday", "Thusday", "Friday", "Saturday"}

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := range 10 {
	// 	fmt.Println(i)
	// }

	// for i := range len(days) {
	// 	fmt.Println(days[i])
	// }

	for _, day := range days {
		if day == "Wednessday" {
			goto shafi
		}
		fmt.Println(day)
	}

shafi:
	fmt.Println("My Name is shafi")
}
