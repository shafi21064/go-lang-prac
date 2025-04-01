package main

import "fmt"

const PRINT_PRICE = 10

type Wall struct {
	width  float32
	height float32
	label  string
}

func main() {

	walls := []Wall{
		{0, 5, "1st wall"},
		{7.5, 5, "2nd wall"},
		{10, 5, "3rd wall"},
	}

	for _, wall := range walls {
		cost, error := calculateWallPrntCost(wall.width, wall.height)
		if error != nil {
			fmt.Println("Error: ", error)
		} else {
			fmt.Printf("%s cost %.2f\n", wall.label, cost)
		}
	}
}

func calculateWallPrntCost(width, height float32) (float32, error) {
	if width <= 0 {
		return 0, fmt.Errorf("Width can't be less or equal zero")
	} else if height <= 0 {
		return 0, fmt.Errorf("height can't be less or equal zero")
	} else {
		area := height * width
		return area * PRINT_PRICE, nil
	}
}
