package practise

import "fmt"

func findLargestElement(arr []int) int {
	largest := arr[0]

	for _, num := range arr {
		if num > largest {
			largest = num
		}
	}

	return largest
}

func FindLargestElement() {
	arr := []int{1, 2, 13, 4, 5}
	largest := findLargestElement(arr)
	fmt.Println("Largest element:", largest)
}
