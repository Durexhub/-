package main

import "fmt"

func evenSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i
		} // обычный цикл для подсчета
	}
	ch <- result // передаем в чан
}
func squareSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i + i
		}
	}
	ch <- result // передаем в чан
}
func main() {
	evenCh := make(chan int) //делаем каналы для наших любимых горутин
	sqCh := make(chan int)   //делаем каналы для наших любимых горутин

	go evenSum(0, 100, evenCh) // создаем горутину GO входные данные + канал для горутниы
	go squareSum(0, 100, sqCh) // создаем горутину GO входные данные + канал для горутниы

	fmt.Println(<-evenCh + <-sqCh) // выводим и подсчет
}
