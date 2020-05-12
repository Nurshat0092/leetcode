package leetcode

func MoveZeros(arr []int) {
	for i, j := 0, 0; j < len(arr); j++ {
		if arr[j] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
}
