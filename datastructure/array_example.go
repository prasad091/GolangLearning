package main

import "fmt"

func main() {

	var arr [5]int
	arr[0] = 100
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println(len(arr2))
}
