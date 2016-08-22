package goutils

const indexNotFound int = -1

// ArrayUtils 用來處理陣列資料
type ArrayUtils struct {
}

// Reverse 用來對 array 變數進行反向排列
func (au ArrayUtils) Reverse(array []byte) []byte {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

// IndexOfUint 用來尋找 array 變數的索引值
func (au ArrayUtils) IndexOfUint(array []uint, valueToFind uint, startIndex int) int {
	if len(array) == 0 {
		return indexNotFound
	}
	if startIndex < 0 {
		startIndex = 0
	}
	for i := startIndex; i < len(array); i++ {
		if valueToFind == array[i] {
			return i
		}
	}
	return indexNotFound
}
