package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// fib returns a function that returns
// successive Fibonacci numbers.
func fib(n big.Int) {
	f1 := big.NewInt(0)
	f2 := big.NewInt(1)

	if n.Cmp(big.NewInt(1)) <= 0 {
		fmt.Println("Ошибка: слишком малое число")
		return
	}

	for range n.Int64() {
		next := big.NewInt(0)
		next.Add(f1, f2)
		f2 = f1
		f1 = next
		fmt.Println("Последовательность: " + f1.String())
	}

}

func main() {
	input := getInput()
	number, err := parseInt(input)
	if err != nil {
		fmt.Printf("Ошибка : %v", err)
	}
	fib(*big.NewInt(int64(number)))
}
func parseInt(input string) (int, error) {
	num1, err := strconv.Atoi(input) // Конвертация в тип int
	if err != nil {
		return 0, fmt.Errorf("введено не число или число слишком велико/мало") // Возвращение ошибки, если не удалась конвертация
	}
	return num1, nil // Возвращение числа
}

func getInput() string {
	fmt.Println("Введите число:")         // Вывод приглашения для пользователя
	scanner := bufio.NewScanner(os.Stdin) // Создание нового сканнера для стандартного ввода
	if scanner.Scan() {                   // Считывание строки из ввода
		return scanner.Text() // Возвращение считанной строки
	}
	if err := scanner.Err(); err != nil { // Проверка на наличие ошибок при сканировании
		fmt.Printf("Ошибка ввода: %v", err)
	}
	return "" // Возвращение пустой строки, если сканнер не смог прочитать ввод
}
