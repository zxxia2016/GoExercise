package main

import (
	"fmt"
	"time"
)

func main10() {

	// fibonacci numbers
	testFibonacci()
	fmt.Println()
	testFibonacci1()
	fmt.Println()
	testMemoizationFibonacci()
	fmt.Println()

	//递归
	ten2One(10)
}

func testFibonacci() {
	for i := 0; i <= 10; i++ {
		var sum, index = fibonacci(i)
		fmt.Println(index, sum)
	}
}
func fibonacci(n int) (int, int) {
	if n <= 1 {
		return 1, n
	}
	var sum, _ = fibonacci(n - 1)
	var sum1, _ = fibonacci(n - 2)
	return sum + sum1, n
}
func fibonacci1() func() int {
	a, b := 1, 1
	return func() int {
		a1 := b
		b = a + b
		a = a1
		return b
	}
}

func testFibonacci1() {
	f := fibonacci1()
	for i := 0; i <= 10; i++ {
		fmt.Println(i, f())
	}
}

func testMemoizationFibonacci() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci2(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

const LIM = 41

var fibs [LIM]uint64

func fibonacci2(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fibs[n] = res
	return
}

func ten2One(n int) {
	fmt.Println(n)
	if n < 1 {
		return
	}
	ten2One(n - 1)
}
