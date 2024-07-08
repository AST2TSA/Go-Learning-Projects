package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// Напишите программу, которая вычисляет факториал введённого пользователем числа.

func main() {
	input := getInput()
	num, err := parseInt(input)
	if err != nil {
		fmt.Printf("Ошибка : %v", err)
		return
	}
	fmt.Println(toFactorial(num))
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
func toFactorial(int1 int) string {
	factorialBody := big.NewInt(1) // Основное тело факториала
	if int1 < 0 {
		return factorialBody.String() // Вовращение единицы, если число является нулём или меньше него.
	}
	for i := range int1 {
		if i < int1 {
			factorialBody.Mul(factorialBody, big.NewInt(int64(i+1))) // Получение факториала
		}
	}
	return factorialBody.String() // Вывод факториала
}