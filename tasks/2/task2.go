package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Напишите программу, которая проверяет, является ли введённое число чётным или нечётным.

func main() {
	input1 := getInput()              // Получение первого числа
	number1, err1 := parseInt(input1) // Конвертация первого числа или получение ошибки
	if err1 != nil {
		fmt.Printf("Ошибка : %v", err1) // Обработка ошибки конвертации
		return
	}
	fmt.Println(isEven(number1))
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

func isEven(number1 int) bool {
	return number1%2 == 0
}
