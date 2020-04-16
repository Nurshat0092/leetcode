package main

import (
	"fmt"

	leetcode ".."
)

func main() {
	arr := []int{1, 5, 2, 0, 6, 8, 0, 6, 0}
	fmt.Println(arr)
	leetcode.DuplicateZeros(arr)
	fmt.Println(arr)
}
