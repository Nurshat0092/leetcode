/*
Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
*/

package leetcode

func SpiralMatrix(matrix [][]int) []int {
	if len(matrix) == 0 || matrix == nil {
		return nil
	}
	N, M := len(matrix), len(matrix[0])
	res := N * M
	direction := 0
	arr := []int{}
	counter := 0
	for {
		switch direction % 4 {
		case 0:
			for i := direction / 4; i < M-direction/4; i++ {
				arr = append(arr, matrix[direction/4][i])
				counter++
			}
		case 1:
			for i := direction/4 + 1; i < N-direction/4; i++ {
				arr = append(arr, matrix[i][M-1-direction/4])
				counter++
			}
		case 2:
			for i := M - 2 - direction/4; i >= direction/4; i-- {
				arr = append(arr, matrix[N-1-direction/4][i])
				counter++
			}
		case 3:
			for i := N - direction/4 - 2; i >= 1+direction/4; i-- {
				arr = append(arr, matrix[i][0+direction/4])
				counter++
			}
		}
		if counter == res {
			return arr
		}
		direction++
	}
	return arr
}
