package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"
	"runtime"
)

func main9() {
	//多个返回值
	v, e := MySqrt(4)
	if e == nil {
		print(v)
	}

	//变长参数
	greeting("hello", "A", "B")
	greetingEverything(1, 1.0, "aaa")

	// defer
	deferTest()
	deferTest2()
	fmt.Println()
	deferFor()
	deferTrace()

	//匿名函数
	testAnonymousFun()
	testDeferAnonymousFun()
}

func MySqrt(a float64) (float64, error) {
	if a < 0 {
		return float64(math.NaN()), errors.New("invalid number")
	}
	return math.Sqrt(a), nil
}

func greeting(prefix string, who ...string) {
	for _, name := range who {
		fmt.Println(prefix, name)
	}
	greeting1(who...)
	greeting2(who)
}

func greeting1(who ...string) {}
func greeting2(who []string)  {}

func greetingEverything(params ...interface{}) {
	for _, p := range params {
		switch p.(type) {
		case string:
			fmt.Println("string", p)
		case int:
			fmt.Println("int", p)
		case float64:
			fmt.Println("float64", p)
		default:
			fmt.Println("default", p)

		}
	}
}

func deferTest() {
	fmt.Println("In function Top")
	defer deferTest1()
	fmt.Println("In function Bottom")
}
func deferTest1() {
	fmt.Println("function will be called in the end")
}

// 结论，只是延迟语句执行时机，变量值是运行到该语句时候的
func deferTest2() {
	for {
		i := 0
		defer fmt.Println(i) // print 0
		i++
		return
	}
}

// 逆序执行（类似栈，即后进先出）print 4,3,2,1,0
// 这么理解：0最先延迟，所以在最后
func deferFor() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func deferTrace() {
	trace(deferTrace)
	defer untrace(deferTrace)
	fmt.Println("deferTrace do something")
}

func trace(i interface{}) {
	fmt.Println(`entering:`, runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name())
	log.Print(`entering:`, runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name())
}

func untrace(i interface{}) {
	fmt.Println(`leaving:`, runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name())
	log.Print(`leaving:`, runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name())
}

func testAnonymousFun() {
	func() {
		fmt.Println(`test anonymous function`)
	}()
}

// print 2
func testDeferAnonymousFun() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

// 闭包的应用：将函数作为返回值
func Add1(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
