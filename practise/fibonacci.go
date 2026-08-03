package practise

import "fmt"

func Fibonacci() {
	n := 10
	fib := []int{0, 1}

	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	fmt.Println("Fibonacci series:", fib)
}
