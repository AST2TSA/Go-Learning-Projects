package main

import (
	"fmt"
	"time"
)

// Напишите программу, которая выводит числа от 10 до 1 с задержкой в одну секунду между выводами.

func main() {
	const secondsToCountDown = 10
	countdown(secondsToCountDown)
}

func countdown(secondsLeft int) {
	const timerSection = 1
	if secondsLeft > 0 {
		fmt.Println(secondsLeft)
		time.Sleep(timerSection * time.Second)
		secondsLeft--
		countdown(secondsLeft)
	} else {
		fmt.Println("Time's up")
	}
}
