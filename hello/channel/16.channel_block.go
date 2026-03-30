package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	go func() {
		time.Sleep(2e9)
		out <- 1
	}()
	println("---------")

	go func() {
		println("---------0")
		time.Sleep(1e9)
		out <- 2
	}()
	sum := <-out
	println(sum)
	sum = <-out
	println(sum)
	println("---------1")

	// select {}
}
