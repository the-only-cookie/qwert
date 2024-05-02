package fibonacci

// FibonacciCalculator структура для вычисления чисел ряда Фибоначчи
type FibonacciCalculator struct{}

// CalculateFibonacciNumber метод для вычисления n-го числа ряда Фибоначчи
func (f *FibonacciCalculator) CalculateFibonacciNumber(n int) int {
	// реализация вычисления числа Фибоначчи
	func (f *Fibonacci) Calculate(n int) int {
		if n <= 1 {
		 return n
		}
	   
		if val, ok := f.cache[n]; ok {
		 return val
		}
	   
		f.cache[n] = f.Calculate(n-1) + f.Calculate(n-2)
		return f.cache[n]
	   }
	   
	   func main() {
		fmt.Println("Введите количество чисел ряда Фибоначчи:")
		var count int
		fmt.Scanln(&count)
	   
		fib := NewFibonacci()
		for i := 0; i < count; i++ {
		 fmt.Printf("%d ", fib.Calculate(i))
		}
}
