/*
Input: nums = [3, 6, 1, 0]
Output: 1
Explanation: 6 is the largest integer, and for every other number in the array x,
6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.

1) nums will have a length in the range [1, 50].
2) Every nums[i] will be an integer in the range [0, 99].
*/

package leetcode

func largetNumberAtLeastTwiceOfOthers(nums []int) int {
	arr := make([]int, 100)
	for i := range nums {
		arr[nums[i]]++
	}
	for i := range nums {
		idx := -1
		for j := 99; j >= nums[i]/2; i-- {
			if arr[j] != 0 {
				idx = i
				break
			}
		}
		if idx != -1 {
			return idx
		}
	}
	return -1
}
