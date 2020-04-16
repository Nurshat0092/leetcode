package main

import (
	"fmt"

	leetcode ".."
)

func main() {
	nums, val := []int{3, 2, 2, 3}, 3
	fmt.Println(leetcode.RemoveElements(nums, val))

}
