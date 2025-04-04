package main

import (
	"fmt"
	"io"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("File writting")
	content := "Hey my name is shafi and currently i am learning go lang"

	filePath := "./output.txt"
	file, err := os.Create(filePath)
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("Length is", length)
	defer file.Close()
	readFile(filePath)
}

func readFile(filePath string) {
	file, err := os.ReadFile(filePath)
	checkError(err)
	fmt.Println("File content: ", string(file))
}
