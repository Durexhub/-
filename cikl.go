package main

import "fmt"

func main() {
	sum_even := 0

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // число нечетное, то пропускаем его
			break    // выход из функции как достигнуто начатое
		} // и переходим к следующей итерации
		sum_even += i
	}
	fmt.Println(sum_even) // выведет: 30

}
