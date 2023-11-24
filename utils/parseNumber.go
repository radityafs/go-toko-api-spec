package utils

import "strconv"

func ParseToNumber(value string) int {
	number, _ := strconv.Atoi(value)
	return number
}
