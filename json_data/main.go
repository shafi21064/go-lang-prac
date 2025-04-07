package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name         string   `json:"name"`
	Price        float32  `json:"price"`
	PlatformName string   `json:"website"`
	Password     string   `json:"-"`
	Tag          []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json in golang")
	EncodeJson()
}

func EncodeJson() {
	myResponse := []Course{
		{Name: "Go lang", Price: 1050, PlatformName: "youtube", Password: "abc123", Tag: []string{"go", "dev"}},
		{Name: "Dart", Price: 599, PlatformName: "Udemy", Password: "bdc123", Tag: []string{"dart", "dev"}},
		{Name: "Flutter", Price: 2050, PlatformName: "coursera", Password: "dcn325", Tag: nil},
	}

	convertedJson, err := json.MarshalIndent(myResponse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", convertedJson)
}
