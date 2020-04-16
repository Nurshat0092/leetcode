package main

import (
	"fmt"

	leetcode ".."
)

func main() {
	nums, nums2 := []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}
	m, n := 3, 3
	fmt.Println(nums)
	leetcode.MergeSortedArray(nums, m, nums2, n)
	fmt.Println(nums)

}
