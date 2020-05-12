package leetcode

func HeightChecker(heights []int) int {
	arr := make([]int, 100)
	res := 0
	for i := range heights {
		arr[heights[i]-1]++
	}
	for i := range heights {
		minPoss, maxPoss := 0, 0
		for j := 0; j < heights[i]-1; j++ {
			minPoss += arr[j]
		}
		maxPoss += minPoss + arr[heights[i]-1]
		if minPoss <= i && i < maxPoss {
			res++
		}
	}
	return res
}
