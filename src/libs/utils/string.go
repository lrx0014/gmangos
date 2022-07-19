package utils

import "strconv"

func ReverseStrings(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
