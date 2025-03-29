package main

import "fmt"

func main() {

	fmt.Println("Wlcome to slice section")

	var myslice = []string{"hey", "hi", "there", "how", "are", "you"}

	fmt.Println("this is my slice", myslice)
	myslice = append(myslice[2:5])
	fmt.Println(myslice)
}
