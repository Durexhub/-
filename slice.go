package main
import "fmt"

func main() {
a := [5]int{0,2,3,4,5} // 1 начинается с двойки
h := a[2:4] // в ней теперь 3 4
}
a := make([]int, 5)
a = append(a, 8) // добаваить в массивы
fmt.Println(a) // выведет [0 0 0 0 0 8]