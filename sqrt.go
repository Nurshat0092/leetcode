package leetcode

func Sqrt(x int) int {
	low, high := 0, x
	for low <= high {
		mid := (low + high) / 2
		mul := mid * mid
		if mul == x {
			return mid
		}
		if mul > x {
			high = mid - 1
			continue
		}
		if mul < x {
			low = mid + 1
			continue
		}
	}
	return high
}
