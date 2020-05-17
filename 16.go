package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	power := 1.0
	if n > 0 {
		power = myPow(x, n/2)
		return power * power * myPow(x, n%2)
	}
}

func main() {
	fmt.Println(myPow(2.0000, -2))
	fmt.Println(myPow(2.0000, 10))
	fmt.Println(myPow(2.10000, 3))
}
