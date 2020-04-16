package leetcode

func SquareOfSortedArray(A []int) []int {
	mid, i, j := 0, 0, 0
	output := []int{}
	for i := range A {
		if A[i] >= 0 {
			mid = i
			break
		}
	}
	i, j = mid, mid-1
	for i < len(A) && j >= 0 {
		if A[i] <= -A[j] {
			output = append(output, A[i]*A[i])
			i++
		} else {
			output = append(output, A[j]*A[j])
			j--
		}
	}
	for i < len(A) {
		output = append(output, A[i]*A[i])
		i++
	}
	for j >= 0 {
		output = append(output, A[j]*A[j])
		j--
	}
	return output
}
