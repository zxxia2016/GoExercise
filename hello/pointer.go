package main

import (
	"fmt"
)

func main6() {
	var i1 = 5
	fmt.Printf("v:%d,%p\n", i1, &i1)

	var str = "11111"
	var pStr *string = &str
	*pStr = "2222"
	fmt.Println(str)
	fmt.Println(*pStr)

}
