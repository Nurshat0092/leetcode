package leetcode

func FindMinimumInRotatedSortedArray(nums []int) int {
	left, right := 0, len(nums)-1
	mid := (right + left) / 2
	for left < right {
		if nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (right + left) / 2
	}
	return nums[mid]
}
