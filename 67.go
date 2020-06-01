package main

import (
	"math"
)

func strToInt(str string) int {
	var result []byte
	var negative bool

	for i := 0; i < len(str); i++ {
		if len(result) == 0 && str[i] == ' ' {
			continue
		}

		if len(result) == 0 && str[i] == '-' {
			negative = true
			continue
		}

		if len(result) == 0 && str[i] == '+' {
			negative = false
			continue
		}

		char := str[i] - '0'
		if len(result) > 0 && char > 9{
			break
		}

		result = append(result, char)
	}

	return rune2int(negative, result)
}

func rune2int(negative bool, data []byte) int {
	var result int

	for i := 0; i < len(data); i++ {
		number := int(data[i])
		base := int(math.Pow10(len(data) - i - 1))
		num := number * base
		result += num
	}

	if negative {
		result = -result
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}

	if result < math.MinInt32 {
		return math.MinInt32
	}

	return result
}

func main() {
	println(strToInt("42"))
	println(strToInt("   -42"))
	println(strToInt("4193 with words"))
	println(strToInt("words and 987"))
	println(strToInt("-91283472332"))
	println(strToInt("+-2"))
}
