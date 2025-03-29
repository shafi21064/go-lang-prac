package main

import (
	"fmt"
)

func main() {
	// var userName string = "shafi"
	// fmt.Println(userName)
	// fmt.Printf("The type of the value is: %T \n", userName)

	// var isLoggedin bool = true
	// fmt.Println(isLoggedin)
	// fmt.Printf("This type of value is: %T \n", isLoggedin)

	var myROll int = 114097
	// fmt.Println(myROll)
	// fmt.Printf("this is a %T type data\n", myROll)
	fmt.Print(myROll)
	// fmt.Printf(myROll);

	/// No var keyword variable
	myName := "Shafi"
	fmt.Println(myName)

	///Same type Muntiple varible in one line
	var s, b, c int = 4, 10, 60
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(c)

	//Group variable declaration
	var (
		roll    int    = 10
		stuName string = "Shafi"
	)
	fmt.Printf("Hello my name is %v and roll is %v", stuName, roll)

	/// Constants
	const PI float64 = 3.1416 //Typed Constants
	fmt.Println(PI)

	const abc = 540 //Untyped Constants
	fmt.Println(abc)

	/// Formatting verbs link
	/// https://www.w3schools.com/go/go_formatting_verbs.php

	/// Arrays
	var myArray = [4]int{5, 10, 25, 25}
	fmt.Printf("This is a array %v \n", myArray[3])

	var mySecondArray = [...]int{5, 10, 20, 25, 30}
	fmt.Printf("This is a non length arrays %v \n", mySecondArray[2])

	myArrayThree := [3]string{"volvo", "hundai", "Hunda"}
	fmt.Println(myArrayThree[2])
}
