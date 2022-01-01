package tools

// HumanizeNumber 10000000 -> 10.000.000
func HumanizeNumber(str string) string {
	var result string
	var count int
	str = ReverseString(str)
	for i, s := range str {
		count++
		result += string(s)
		if i == len(str)-1 || string(str[i+1]) == `-` { // don't set last dot
			continue
		}
		if count == 3 {
			count = 0
			result += "."
		}
	}
	return ReverseString(result)
}
