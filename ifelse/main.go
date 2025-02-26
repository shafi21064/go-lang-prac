package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("if else")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	number, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	fmt.Println(number)

	if number > 10 {
		fmt.Println("Less than 10")
	} else if number < 10 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("It is 10")
	}

	// if newNumber := 20; newNumber < 10 {
	// 	fmt.Println("print")
	// }

}
