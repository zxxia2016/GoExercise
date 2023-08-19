package main

import "fmt"

func main4() {
	testSwitch(100)
}

func calculateSummary(a int, b int) int {
	return a + b
}

func testSwitch(condition int) {
	switch condition {
	case 1, 11:
		fmt.Println("case 1")
		fmt.Println("case 11")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("other condition")
	}

	switch {
	case condition > 1:
		fmt.Println(`condition greater than 1`)
	case condition == 1:
		fmt.Println("condition equal to 1")
	default:
		fmt.Println("condition smaller than 1")
	}

	switch d := calculateSummary(1, 2); {
	case d > 1:
		fmt.Print("111111")
	}
}
