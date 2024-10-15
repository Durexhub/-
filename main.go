package main

import (
	"fmt"
)

func main() {
	var (
		age  int
		name string
	)
	fmt.Scan(&age, &name)
	fmt.Printf("меня зовут %s\n мне столько то лет %d", name, age)
}
