package main

import "fmt"

func main() {
	height := 177
	if height <= 165 {
		fmt.Println("Человек невысокого роста")
	} else if height > 165 && height <= 185 {
		fmt.Println("Человек среднего роста")
	} else {
		fmt.Println("Высокий человек")
	}
}
