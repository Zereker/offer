package main

import "fmt"

func minArray(numbers []int) int {
	if numbers == nil {
		return 0
	}

	minimum := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if minimum > numbers[i] {
			minimum = numbers[i]
		}
	}

	return minimum
}

func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
}
