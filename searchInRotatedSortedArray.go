package leetcode

func SearchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 || nums == nil {
		return -1
	}
	atStart := target >= nums[0]
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high + low) / 2
		if nums[mid] == target {
			return mid
		}
		if atStart {
			atWrong := nums[0] > nums[mid]
			if nums[mid] > target || atWrong {
				high = mid - 1
				continue
			}
			if nums[mid] < target {
				low = mid + 1
				continue
			}
		}
		if !atStart {
			atWrong := nums[len(nums)-1] < nums[mid]
			if nums[mid] < target || atWrong {
				low = mid + 1
				continue
			}
			if nums[mid] > target {
				high = mid - 1
				continue
			}
		}
	}
	return -1
}
