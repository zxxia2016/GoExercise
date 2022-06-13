// $ go mod init example/hello
// $ go run .
// $ go help
// $ go mod tidy
// go: finding module for package rsc.io/quote
// go: found rsc.io/quote in rsc.io/quote v1.5.2
// $ go env -w GOPROXY=https://goproxy.cn,direct 设置代理，否则包无法下载
// search for a "quote" package. https://pkg.go.dev/search?q=quote

// TODO: https://go.dev/doc/tutorial/create-module

package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}