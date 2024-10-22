package main

import "fmt"

func main() {
	a := make([]int, 5)
	a[1] = 2              // присвоение числа массиву
	for i, v := range a { // вычисление индекса и значение массива можно пропустить и вычислить _,
		fmt.Print(i + v)
	}
}
