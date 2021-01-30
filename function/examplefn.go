package main

import (
	"fmt"
)

func add(x int, y int) int {
	var out = x + y
	return out
}

func main() {
	num1 := 5
	num2 := 4

	result := add(num1, num2)
	fmt.Println(result)

}
