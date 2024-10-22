package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["james"] = 42 // присваиваем значения мапе по словам можно иначе
	m["lel"] = 22
	m["heh"] = 21
	h := map[int]string{ // без make и более удобно присвоение
		22: "james",
		31: "fgsdfg"}
	fmt.Print(h)
}
func rain() {
	sum := 0
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"} // входящие
	var men string                                                                            //входной последний матч
	fmt.Scanln(&men)
	results = append(results, men) // добавили последний матч
	h := make(map[string]int)      // обозначаем вход выход с карты
	h["w"] = 3
	h["d"] = 1
	h["l"] = 0
	for _, x := range results { // плюсуем ренджем весь массив
		sum += h[x]
	}
	fmt.Print(sum) // главно не в цикл
}
