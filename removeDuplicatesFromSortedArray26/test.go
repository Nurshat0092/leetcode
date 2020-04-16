package main

import (
	"fmt"

	leetcode ".."
)

func main() {
	nums1 := []int{1, 1, 2}
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Input: ", nums1)
	fmt.Println("Output: ", leetcode.RemoveDuplicatesFromSortedArray(nums1))
	fmt.Println("*******")
	fmt.Println("Input: ", nums2)
	fmt.Println("Output: ", leetcode.RemoveDuplicatesFromSortedArray(nums2))
}
