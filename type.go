package main

import "fmt"

func main() {
	type Age struct { // type - тип, Название с большой буквы и stuct
		name string
		age  int
	}
	ph := &Age{"Lel", 20} //Присвоение указателя но можно и заебаться ниже
	p := Age{"l", 2}      // таким образом сделали в 2 шага но проще запомнить первый
	a := &p
	fmt.Print(a.age, ph.name)

}
