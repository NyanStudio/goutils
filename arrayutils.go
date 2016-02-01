package goutils

type ArrayUtils struct {

}

func (au ArrayUtils) Reverse(array []byte) []byte {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}
