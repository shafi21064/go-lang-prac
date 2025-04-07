package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL string = "http://localhost:8000/get"

func main() {
	fmt.Println("Handel web server request")
	PerformGetRequest()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformGetRequest() {
	response, err := http.Get(URL)
	CheckError(err)
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	content, err := io.ReadAll(response.Body)

	fmt.Println("Content: ", content)

	fmt.Println(string(content))
}
