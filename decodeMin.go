package main

import (
	"fmt"
	"strings"
)

func decodeToMinNumber(symbols string) string {
	n := len(symbols) + 1
	result := make([]int, n)

	for i := range result {
		result[i] = 0
	}

	for i := 0; i < len(symbols); i++ {
		switch symbols[i] {
		case 'L':
			if result[i] <= result[i+1] {
				result[i] = result[i+1] + 1
			}
		case 'R':
			if result[i] >= result[i+1] {
				result[i+1] = result[i] + 1
			}
		case '=':
			if result[i] != result[i+1] {
				result[i+1] = result[i]
			}
		}
	}

	for i := 0; i < len(symbols); i++ {
		if symbols[i] == 'L' && result[i] <= result[i+1] {
			result[i] = result[i+1] + 1
		}
	}

	if strings.HasPrefix(symbols, "=") {
		for i := 0; i < len(symbols)-1; i++ {
			if symbols[i] == '=' {
				if result[i] < result[i+1] {
					result[i] = result[i+1]
				}
			}
		}
	}

	var sb strings.Builder
	for _, v := range result {
		sb.WriteString(fmt.Sprintf("%d", v))
	}
	return sb.String()
}

func decode() {
	inputs := []string{"LLRR=", "==RLL", "=LLRR", "RRL=R"}
	for _, input := range inputs {
		output := decodeToMinNumber(input)
		fmt.Printf("Input: %s, Output: %s\n", input, output)
	}

}
