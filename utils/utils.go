package utils

import (
	"strconv"
	"strings"
)

func MaskHalfInt(input int) string {
	return MaskHalf(strconv.Itoa(input))
}

func MaskHalfInt64(input int64) string {
	return MaskHalf(strconv.FormatInt(input, 10))
}

func MaskHalf(input string) string {
	if input == "" {
		return input
	}
	if len(input) < 2 {
		return input
	}
	length := len(input)
	visibleLength := length / 2
	maskedLength := length - visibleLength
	return input[:visibleLength] + strings.Repeat("*", maskedLength)
}
