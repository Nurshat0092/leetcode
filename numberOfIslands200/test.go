package main

import (
	"fmt"

	leetcode ".."
)

func main() {
	island := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(leetcode.NumberOfIslands(island))
}
