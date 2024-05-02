package console

import (
	"fmt"
	"math"
)
import  "fibonacci"
import "utils"


// FibonacciConsoleApp функция для консольного приложения
func FibonacciConsoleApp() {
	fibonacciCalculator := &fibonacci.FibonacciCalculator{}
	closestFibonacciNumberFinder := &utils.ClosestFibonacciNumberFinder{}

	var inputNumber int
	fmt.Print("Введите число: ")
	fmt.Scan(&inputNumber)

	if isFibonacciNumber(inputNumber, fibonacciCalculator) {
		fmt.Println("Введенное число принадлежит ряду Фибоначчи.")
	} else {
		closestFibonacciNumber := closestFibonacciNumberFinder.FindClosestFibonacciNumber(inputNumber)
		fmt.Println("Ближайшее число из ряда Фибоначчи: ", closestFibonacciNumber)
	}
}

// isFibonacciNumber функция для проверки, является ли число числом ряда Фибоначчи
func isFibonacciNumber(number int, fibonacciCalculator *fibonacci.FibonacciCalculator) bool {
	// реализация проверки принадлежности числа к ряду Фибоначчи

	// Проверяем, является ли выражение 5*n*n+4 или 5*n*n-4 полным квадратом
	return isPerfectSquare(5*number*number+4) || isPerfectSquare(5*number*number-4)
}

// isPerfectSquare функция для проверки, является ли число полным квадратом
func isPerfectSquare(n int) bool {
	root := int(math.Sqrt(float64(n)))
	return root*root == n
}
