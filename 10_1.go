package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return memoize(n)
}

func memoize(n int) int {
	cache := map[int]int{0: 0, 1: 1}

	for i := 2; i <= n; i++ {
		cache[i] = (cache[i-1] % (1e9 + 7)) + (cache[i-2] % (1e9 + 7))
	}

	return cache[n] % (1e9 + 7)
}

func main() {
	fmt.Println(fib(45))
}
