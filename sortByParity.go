package leetcode

func SortByParity(A []int) []int {
	for i, j := 0, 0; j < len(A); j++ {
		if A[j]%2 == 0 {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	return A
}
