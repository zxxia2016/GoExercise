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

// demo v4: Return an random greetings

// demo v5: Return greetings for multiple people

// demo v6: Add a test
// Ending a file's name with _test.go tells the go test command that this file contains test functions.
// $ go test (-v)
// Example: greetings_test.go

// demo v7: Compile and install the application
// $ go build // >> hello.exe
// $ go list -f '{{.Target}}' // discover the install path
// $ go env -w GOBIN=C:\path\to\your\bin // set install path
// $ go install

// TODO:
// https://go.dev/doc/tutorial/create-module
// https://go.dev/doc/effective_go
// A Tour of Go: https://go.dev/tour
// Managing dependencies: https://go.dev/doc/modules/managing-dependencies
// Developing and publishing modules: https://go.dev/doc/modules/developing

package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	// demo v1 Hello World
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// demo v2: Create a Go module
	message1 := greetings.Hello1("Gladys")
	fmt.Println(message1)

	// demo v3: Return and handle an error
	message2, err := greetings.Hello2("Gladys")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message2)

	// demo v4: Return an random greetings
	message3, err := greetings.Hello3("Gladys")
	fmt.Println(message3)

	// demo v5: Return greetings for multiple people
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
