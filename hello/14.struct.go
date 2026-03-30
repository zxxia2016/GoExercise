package main

import "fmt"

type A struct{ a int }
type B struct{ a, b int }

type C struct {
	A
	B
}

var c C

type D struct {
	B
	b float32
}

var d D

func main14() {
	// ambiguous DOT reference c.a disambiguate with either c.A.a or c.B.a。
	// fmt.Println(c.a)
	fmt.Println(d.b)
	fmt.Println(d.B.b)
}
