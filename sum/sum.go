package sum

func Sum(arr []int) int {
	result := 0
	for _, i := range arr {
		result += i
	}
	return result
}
