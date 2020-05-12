package leetcode

func FindPeakElement(nums []int) int {
	left, high := 0, len(nums)-1
	mid := (left + high) / 2
	for left < high {
		if nums[mid] > nums[mid+1] {
			high = mid
		} else if nums[mid] < nums[mid+1] {
			left = mid + 1
		}
		mid = (left + high) / 2
	}
	return mid
}
