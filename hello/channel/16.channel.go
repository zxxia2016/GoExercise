package main

import (
	"fmt"
	"runtime"
	"time"
)

func main16() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan string)
	go sendData(ch)
	go recvData(ch)

	for {
		time.Sleep(2 * time.Second)
		PrintLog("sleep")
	}
}

func sendData(ch chan string) {
	PrintLog("send hello")
	ch <- "hello"
	//阻塞，等一条取完再执行这代码
	// runtime.Goexit() //让协程退出
	PrintLog("send world")
	ch <- "world"
}
func recvData(ch chan string) {
	time1 := time.NewTimer(1 * time.Second)
	<-time1.C
	for {
		PrintLog("get ", <-ch)
	}
}

func PrintLog(s ...string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), s)
}

func numEchoRange(in <-chan string, done chan<- bool) {
	for num := range in {
		fmt.Printf("%s\n", num)
	}
	done <- true
}
