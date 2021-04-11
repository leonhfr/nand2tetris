package util

import (
	"fmt"
	"strings"
)

func DecimalToBinary(decimal int) string {
	var binary []int

	for decimal != 0 {
		binary = append(binary, decimal%2)
		decimal = decimal / 2
	}

	if len(binary) == 0 {
		return "0"
	}

	var output string
	for i := len(binary) - 1; i >= 0; i-- {
		output += fmt.Sprint(binary[i])
	}
	return output
}

func PadZeroLeft(str string, length int) string {
	pad := length - len(str)
	if pad <= 0 {
		return str
	}
	return strings.Repeat("0", pad) + str
}
