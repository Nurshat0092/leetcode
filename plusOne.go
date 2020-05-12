package leetcode

func PlusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 1; i-- {
		if digits[i] > 9 {
			digits[i-1]++
		}
		digits[i] %= 10
	}
	if digits[0] > 9 {
		digits[0] %= 10
		digits = append([]int{1}, digits...)
	}
	return digits
}
