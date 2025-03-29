package main

import "fmt"

func main() {
	fmt.Println("Structs")

	shafi := User{Name: "shafi", Email: "shafi@go.dev", Status: true, Age: 23}
	fmt.Println(shafi)
	fmt.Printf("Details are %+v", shafi)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
