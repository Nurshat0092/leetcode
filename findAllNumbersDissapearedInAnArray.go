package leetcode

func FindAllNumbersDissapearedInAnArray(nums []int) []int {
	res := []int{}
	for i := range nums {
		if nums[abs(nums[i])-1] > 0 {
			nums[abs(nums[i])-1] = -nums[abs(nums[i])-1]
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(num1 int) int {
	if num1 < 0 {
		return -num1
	}
	return num1
}
