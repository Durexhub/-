package main

import "fmt"

func change(x int) { // дефолтное значение и присвоение через нам будет 14
	x = 22
}

func changePointer(pointer *int) { //при присвоении преобладает и будет 44
	*pointer *= 22 // 44
}

func main() {
	num := 14

	change(num)
	fmt.Println(num) // выведет 14

	changePointer(&num) //& - служит указателем  на ячейку
	fmt.Println(num)    // выведет 22
	type Age struct {
		name string
		age  int
	}
	x := Age{"an", 18}
	fmt.Print(Name: x.age, x.name)

}
