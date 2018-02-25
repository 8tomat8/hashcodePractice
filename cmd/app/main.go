package app

var r = 100
var c = 100
var tomatoes = true
var mushrooms = false

func maxSliceCount(input [r][c]bool) int {
	t := 0
	arraySize := r * c
	for _, v := range input {
		for _, v1 := range v {
			if v1 {
				t++
			}
		}
	}
	if arraySize-t > t {
		return t
	}
	return arraySize - t
}