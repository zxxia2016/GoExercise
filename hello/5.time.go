package main

import (
	"fmt"
	"time"
)

func main5() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%02d", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println(t)

	test := 1e9 //10的9次方
	fmt.Println(test)
	week := 60 * 60 * 24 * 7 * 1e9
	weekFromNow := t.Add(time.Duration(week))
	fmt.Println(weekFromNow)

	start := time.Now()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("delta:", delta)
}
