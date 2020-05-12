package main

import (
	"fmt"

	leetcode ".."
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(leetcode.SearchInRotatedSortedArray(nums, target))
}
