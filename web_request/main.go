package main

import (
	"fmt"
	"io"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

const URL = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	response, err := http.Get(URL)
	checkError(err)

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	checkError(err)
	fmt.Println(string(data))
}
