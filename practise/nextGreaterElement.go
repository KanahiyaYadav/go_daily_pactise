package practise

import "fmt"

func NextGreaterElement() {
	arr := []int{1, 3, 2, 4}
	var newArr []int

	for i := 0; i < len(arr); i++ {
		nextGreater := -1
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				nextGreater = arr[j]
				break
			}
		}

		newArr = append(newArr, nextGreater)
	}

	fmt.Println(newArr)
}
