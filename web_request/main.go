package main

import (
	"fmt"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

const URL = "https://lco.dev"

func main() {
	response, err := http.Get(URL)
	checkError(err)
	fmt.Println(response)
}
