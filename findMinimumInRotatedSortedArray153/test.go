package main

import (
	"fmt"

	leetcode ".."
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(leetcode.FindMinimumInRotatedSortedArray(nums))
}
