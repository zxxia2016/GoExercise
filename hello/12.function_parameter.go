package main

import (
	"fmt"
	"strings"
)

func main12() {
	//函数作为参数
	callback(1, Add)
	testIndexFunc()
	myStrReplace()
}

func callback(a int, f func(int, int)) {
	f(a, 2)
}

func Add(a, b int) {
	fmt.Println(a + b)
}

func indexFunc(a rune) bool {
	if a > 255 {
		return true
	}
	return false
}

func testIndexFunc() {
	index := strings.IndexFunc("abc,日报", indexFunc)
	fmt.Println(`testIndexFunc:`, index)
}

func myStrReplace() {
	onlyAscii := func(c rune) rune {
		if c > 127 {
			return ' '
		}
		return c
	}
	fmt.Println(strings.Map(onlyAscii, "abc我是谁d"))
}
