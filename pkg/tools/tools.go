package tools

func ContainNumberInSlice(number int, slice []int) bool {
	for _, elem := range slice {
		if number == elem {
			return true
		}
	}
	return false
}
