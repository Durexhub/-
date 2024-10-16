package main

import (
	"fmt"
)

func getFruitColor(fruit string) string {
	switch fruit {
	case "apple":
		return "Red"
	case "banana":
		return "Yellow"
	case "orange":
		return "Orange"
	default:
		return "Unknown fruit"
	}
}

func main() {
	color := getFruitColor("banana")
	fmt.Println("The color of the banana is:", color)

	color = getFruitColor("apple")
	fmt.Println("The color of the apple is:", color)

	color = getFruitColor("grape")
	fmt.Println("The color of the grape is:", color)
}
