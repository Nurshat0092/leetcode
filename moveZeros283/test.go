package main

import (
	"fmt"

	leetcode ".."
)

func main() {
	arr := []int{0, 0, 1, 4, 6, 0, 12, 13, 4, 0}
	fmt.Println(arr)
	leetcode.MoveZeros(arr)
	fmt.Println(arr)
}
