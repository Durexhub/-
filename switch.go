package main

import "fmt"

func main() {
	defer ifa(0)
	c := ifa(80) // после возврата функции нужно указать значение которые будут в ней например ifa("apple")
	switch {     // конкретные условия как if else
	case c == 24:
		fmt.Printf("like%d", c)
		fallthrough // если нужна вторая версия - переносит на кейс ниже, ставим в каждом кейсе до
	case c >= 24:
		fmt.Print("ровно")
	default:
		fmt.Print("промокод неверный")
	}
}
func ifa(a int) int {
	switch a { // почти как if но для поиска условно по номерам какая цифра в переменной такой кейс выдаст
	case 80:
		return a * 2
	default: // обязетльно иначе коду некуда бежать дальше и работать ничего не будет
		return 0
	}
}
