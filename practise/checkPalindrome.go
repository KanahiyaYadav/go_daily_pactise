package practise

import (
	"fmt"
	"strings"
)

func reverseString1(s string) string {
	runes := []rune(s)

	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func CheckPalindrome() {
	s := "Naman"
	rs := reverseString1(s)
	lowerS := strings.ToLower(s)
	lowerRS := strings.ToLower(rs)
	if lowerS == lowerRS {
		fmt.Println("String is palindrome")
	} else {
		fmt.Println("String is not palindrome")
	}

}
