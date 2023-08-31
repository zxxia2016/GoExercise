package main

import "fmt"

type Vector3D [3]float32

var vec Vector3D

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

var SCREEN [WIDTH][HEIGHT]pixel

func main() {

	var array [5]int
	// 数组赋值
	for i := 0; i < len(array); i++ {
		array[i] = i + 1
	}
	// 数组赋值
	a := [...]string{"a", "b", "c", "d"}
	fmt.Println(a) //[a b c d]
	// 数组初始化赋值赋值
	b := []int{1, 2, 3}
	fmt.Println(b) //[1 2 3]

	//遍历
	for i, v := range array {
		fmt.Println(i, v)
	}

	// 直接赋值是引用
	c := b
	// 传参是引用=实参
	changeArray(c) //[5 2 3]
	fmt.Println(b) //[5 2 3]

	// 数组指针传参
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}

	setScreenPixel()
}

func fp(a *[3]int) { fmt.Println(a) }

func changeArray(a []int) {
	a[0] = 5
	fmt.Println(a)
}
func setScreenPixel() {
	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			SCREEN[i][j] = 0
		}
	}
	// fmt.Println(SCREEN)
}
