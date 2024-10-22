package main

import "fmt"

func main() {
	// Исходные результаты матчей
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	// Ожидаемый результат последнего матча (ввод от пользователя)
	var lastResult string
	fmt.Print("Введите результат последнего матча (w/l/d): ")
	fmt.Scanln(&lastResult) // Чтение ввода пользователя

	// Добавление результата последнего матча к массиву
	results = append(results, lastResult)

	// Создание карты для хранения очков за каждый результат
	pointsMap := map[string]int{
		"w": 3, // победа
		"d": 1, // ничья
		"l": 0, // проигрыш
	}

	// Переменная для подсчета очков
	totalPoints := 0

	// Подсчет очков
	for _, result := range results {
		totalPoints += pointsMap[result] // Получаем очки из карты
	}

	// Вывод общего количества очков
	fmt.Printf("Общее количество очков: %d\n", totalPoints)
}
