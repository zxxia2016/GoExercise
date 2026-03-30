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

func main13() {

	array := [5]int{}
	// 数组赋值
	for i := 0; i < len(array); i++ {
		array[i] = i + 1
	}

	// 数组赋值
	a := [...]string{"a", "b", "c", "d"}
	fmt.Println(a) //[a b c d]

	//数组拷贝
	array2 := array
	array2[0] = 0
	fmt.Println(array)  //[1 2 3 4 5]
	fmt.Println(array2) //[0 2 3 4 5]

	//遍历
	for i, v := range array {
		fmt.Println(i, v)
	}

	// 切片初始化赋值
	b := []int{1, 2, 3}
	fmt.Println(b) //[1 2 3]

	// 直接赋值是引用
	c := b
	// 传参是引用=实参
	changeArray(c) //[5 2 3]
	fmt.Println(b) //[5 2 3]

	// 数组指针传参
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}

	// 二位数组
	setScreenPixel()

	makeslice := make([]int, 5)
	l := len(makeslice)
	for i := 0; i < l; i++ {
		makeslice[i] = i
	}

	//切片追加
	makeslice = append(makeslice, 5, 6, 7)
	//切片取值
	referslice := makeslice[1:3]
	fmt.Println(referslice)

	//切片复制
	copyslice := make([]int, len(makeslice))
	n := copy(copyslice, makeslice)
	copyslice[0] = 1
	fmt.Println(makeslice)                //[0 1 2 3 4 5 6 7]
	fmt.Println(copyslice)                //[1 1 2 3 4 5 6 7]
	fmt.Printf("Copied %d elements\n", n) // n == 5

	s := "\u00ff\u745c"
	for i, v := range s {
		fmt.Printf("i: %d,v: %c\n", i, v)
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("i: %d,v: %c\n", i, s[i])
	}
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
