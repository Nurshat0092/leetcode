package leetcode

func ReplaceElementsWithGreatestElementOnRightSide(arr []int) []int {
	maxNum := arr[len(arr)-1]
	prev := arr[len(arr)-1]

	for i := len(arr) - 2; i >= 0; i-- {
		if maxNum < prev {
			maxNum = prev
		}
		prev = arr[i]
		arr[i] = maxNum
	}
	arr[len(arr)-1] = -1
	return arr
}
