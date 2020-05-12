package leetcode

func ValidMountainArray(A []int) bool {
	i, j := 0, len(A)
	for i+1 < j && A[i] < A[i+1] {
		i++
	}
	if i == 0 || i == j-1 {
		return false
	}
	for i+1 < j && A[i] > A[i+1] {
		i++
	}
	return i == j-1
}
