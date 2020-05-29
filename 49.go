package main

func min(num1, num2, num3 int) int {
	if num1 <= num2 {
		if num1 <= num3 {
			return num1
		}
	} else {
		if num2 <= num3 {
			return num2
		}
	}

	return num3
}

func nthUglyNumber(n int) int {
	if n <= 0 {
		return -1
	}

	p2, p3, p5 := 0, 0, 0

	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		dp[i] = min(dp[p2]*2, dp[p3]*3, dp[p5]*5)

		if dp[i] == dp[p2]*2 {
			p2++
		}

		if dp[i] == dp[p3]*3 {
			p3++
		}

		if dp[i] == dp[p5]*5 {
			p5++
		}
	}

	return dp[n-1]
}

func main() {
	println(nthUglyNumber(10))
}
