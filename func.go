package main

import (
	"fmt"
)

func main() {
	jet := concat("y", "ra")      //ВОЗВРАЩАЕТ ЗНАЧЕНИЕ ПОСЧИТАННОЕ В КАКОЙ ЛИБО  ФУНКЦИИ И ДАЛЬШЕ С ЭТИМ ЗНАЧЕНИЕМ МОЖНО РАБОТАТЬ
	bet := concat1("pri", "vet ") //В ДРУГИХ ФУКНЦИЯХ
	let := bet + jet
	fmt.Printf("%s", let)
}
func concat(be, ve string) string { //ИЗНАЧАЛЬНО УКАЗЫВАЕМ ВХОДНЫЕ ЧИСЛА - (be, ve string) ПОСЛЕ RETURN ИДУТ ВЫХОДНЫЕ ИХ УКАЗЫВАЕМ
	return be + ve // ЛИБО INT ЛИБО В СКОБКАХ (int,string) вспомнить магазин
} // условно где то в поиск вбили дальше эти значения идут в другую функцию все так просто спасибо Юра из чата
func concat1(be, ve string) string {
	return be + ve
}
