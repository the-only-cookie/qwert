package utils

import "math"

// ClosestFibonacciNumberFinder структура для поиска ближайшего числа из ряда Фибоначчи
type ClosestFibonacciNumberFinder struct{}

// FindClosestFibonacciNumber метод для нахождения ближайшего числа из ряда Фибоначчи к заданному числу
func (c *ClosestFibonacciNumberFinder) FindClosestFibonacciNumber(target int) int {
	// реализация поиска ближайшего числа из ряда Фибоначчи

	if target <= 0 {
		return 0
	}

	fibPrev := 0
	fibCurr := 1
	fibNext := fibPrev + fibCurr

	for fibNext <= target {
		fibPrev = fibCurr
		fibCurr = fibNext
		fibNext = fibPrev + fibCurr
	}

	if math.Abs(float64(fibCurr-target)) < math.Abs(float64(fibNext-target)) {
		return fibCurr
	}
	return fibNext
}
