package leetcode

/*
Input:
nums = [1, 7, 3, 6, 5, 6]
Output: 3
Explanation:
The sum of the numbers to the left of index 3 (nums[3] = 6) is equal to the sum of numbers to the right of index 3.
Also, 3 is the first index where this occurs.
*/
func FindPivotIndex(nums []int) int {
	sum, curSum := 0, 0
	for i := range nums {
		sum += nums[i]
	}
	for i := range nums {
		if curSum == sum-curSum-nums[i] {
			return i
		}
		curSum += nums[i]
	}
	return -1
}
