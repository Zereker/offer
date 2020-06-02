package main

func printNumbers(n int) []int {
	max := 1
	for i := 0; i < n; i++ {
		max *= 10
	}

	result := make([]int, 0, max-1)
	for i := 1; i < max; i++ {
		result = append(result, i)
	}

	return result
}

func main() {
	println(printNumbers(2))
}
