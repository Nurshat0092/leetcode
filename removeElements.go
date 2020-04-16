package leetcode

func RemoveElements(arr []int, val int) int {
	i, j := 0, len(arr)
	for i < j {
		if arr[i] == val {
			arr[i] = arr[j-1]
			j--
		} else {
			i++
		}
	}
	return j
}
