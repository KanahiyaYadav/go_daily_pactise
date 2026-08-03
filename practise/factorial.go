package practise

import "fmt"

func factorialUsingLoop(n int) int {
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial = factorial * i
	}

	return factorial

}

func factorialUsingRecursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialUsingRecursion(n-1)
}

func FactorialOfNumber() {
	n := 5

	factorial := factorialUsingLoop(n)
	fmt.Println("Factorial of", n, "is", factorial, "using loop")

	factorial = factorialUsingRecursion(n)
	fmt.Println("Factorial of", n, "is", factorial, "using recursion")

}
