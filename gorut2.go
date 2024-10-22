package main

import "fmt"

//определите функцию
func download(a int, ch chan int) { // входное число плюс передача числа обратно в горутину
	nums := 0
	for f := a; f >= 0; f-- { // цикл для подсчета числа
		nums += f
	}
	ch <- nums // передача числа в чан
}

func main() {
	ch1 := make(chan int) // каналы для горутин
	ch2 := make(chan int) // каналы для горутин
	ch3 := make(chan int) // каналы для горутин

	var s1, s2, s3 int
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Scanln(&s3)

	go download(s1, ch1) //cоздание горутин отправление аргумента и сбор
	go download(s2, ch2) //cоздание горутин отправление аргумента и сбор
	go download(s3, ch3) //cоздание горутин отправление аргумента и сбор
	fmt.Print(<-ch1 + <-ch2 + <-ch3)
	//выведите сумму всех результатов
}
