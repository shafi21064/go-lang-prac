package main

import (
	"fmt"
	"time"
)

func main() {
	// var userName string = "shafi"
	// fmt.Println(userName)
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Printf("Input a number: ")

	// input, _ := reader.ReadString('\n')
	// fmt.Println("your number is", input)
	// fmt.Printf("your number is a %T \n", input)
	// convertedNumber, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("added number is ", convertedNumber+1)
	// }

	currentTIme := time.Now()

	fmt.Println("current date is ", currentTIme.Format("01-02-2006 Monday"))
}
