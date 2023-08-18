package main

import "fmt"

func main8() {
	// 习题5.5
	for i := 0; i < 6; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}

	str := "G"
	for i := 0; i < 6; i++ {
		fmt.Println(str)
		str += "G"
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}

	// for&switch
	for i := 0; i < 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

	fmt.Println()

	// for 替换while
	i := 1
	for i >= 0 {
		i--
		fmt.Println(i)
	}

	// for break continue
	i = 3
	for {
		i--
		if i < 0 {
			break
		}
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println()

	// for-range
	str = "I am a student."
	for i, char := range str {
		fmt.Printf("%d, %c\n", i, char)
	}

	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}

	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n",
			index, rune, rune, rune, []byte(string(rune)))
	}

	fmt.Println()

	// GOTO 用法用于跳出错误
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	fmt.Println()

	i = 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
