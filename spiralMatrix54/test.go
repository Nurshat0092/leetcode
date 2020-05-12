package main

import (
	"fmt"

	leetcode ".."
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(leetcode.SpiralMatrix(matrix))
}
