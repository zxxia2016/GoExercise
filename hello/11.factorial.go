package main

import (
	"fmt"
)

func main11() {
	for i := 0; i < 10; i++ {
		result := factorial(i)
		fmt.Println(i, result)
	}
}

// 阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
