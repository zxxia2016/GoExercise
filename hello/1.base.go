package main

import (
	"fmt"
	"log"
	"runtime"
)

func main1() {
	a, b := 1, 2
	// 这3行等价与下面那一行，说明运算符从右到左执行
	// a, b = b, a+b
	a1 := b
	b = a + b
	a = a1
	fmt.Println(a, b)

	// 输出调用文件名和函数
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
}
