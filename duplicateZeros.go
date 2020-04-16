package leetcode

func DuplicateZeros(arr []int) {
	duplNum := 0
	size := len(arr) - 1
	for i := 0; i <= size-duplNum; i++ {
		if arr[i] == 0 {
			if i == size-duplNum {
				arr[size] = 0
				size--
				break
			}
			duplNum++
		}
	}
	last := size - duplNum
	for i := last; i >= 0; i-- {
		if arr[last] == 0 {
			arr[last+duplNum] = 0
			duplNum--
			arr[last+duplNum] = 0
		} else {
			arr[last+duplNum] = arr[i]
		}
	}
}

// [1 5 2 0 6 8 0 6]
