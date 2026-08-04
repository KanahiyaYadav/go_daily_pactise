package practise

import "fmt"

func isNumberPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}

func PrimeNumberCheck() {
	num := 2

	if isNumberPrime(num) {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}
}
