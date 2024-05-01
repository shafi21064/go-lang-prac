package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("User Input File")

	fmt.Println("Enter your Name: ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Printf("Thank you %v for your response", input)
}
