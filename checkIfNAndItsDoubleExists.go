package leetcode

func CheckIfNAndItsDoubleExists(arr []int) bool {
	mapG := map[int]bool{}
	for i := range arr {
		if mapG[arr[i]*2] || (arr[i]%2 == 0 && mapG[arr[i]/2]) {
			return true
		}
		mapG[arr[i]] = true
	}
	return false
}
