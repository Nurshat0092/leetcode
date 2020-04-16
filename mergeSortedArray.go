package leetcode

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}
	for j >= 0 {
		nums1[j] = nums2[j]
		j--
	}
}

/*
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/
