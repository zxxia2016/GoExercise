// demo v1 Hello World

// $ go mod init example/hello
// $ go run .
// $ go help
// $ go mod tidy
// go: finding module for package rsc.io/quote
// go: found rsc.io/quote in rsc.io/quote v1.5.2
// $ go env -w GOPROXY=https://goproxy.cn,direct 设置代理，否则包无法下载
// search for a "quote" package. https://pkg.go.dev/search?q=quote

// demo v2: Create a Go module

// $ go mod edit -replace example.com/greetings=../greetings 替换本地依赖
// $ go mod tidy 同步依赖

// demo v3: Return and handle an error

// TODO: https://go.dev/doc/tutorial/create-module

package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	message, errors := greetings.Hello("")
	if errors != nil {
		log.Fatal(errors)
	}
	fmt.Println(message)
}
