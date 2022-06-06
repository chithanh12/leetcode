package day11

import (
	"strings"
)

var (
	numbers = map[byte]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}
)

func letterCasePermutation(s string) []string {
	return permutation([]byte(s))
}

func permutation(inputs []byte) []string {
	if len(inputs) == 1 {
		return possibleValues(inputs[0])
	}

	tmp := permutation(inputs[:len(inputs)-1])
	pos := possibleValues(inputs[len(inputs)-1])
	rs := make([]string, 0)
	for idx, _ := range tmp {
		for idx1, _ := range pos {
			xx := tmp[idx] + pos[idx1]
			rs = append(rs, string(xx))
		}
	}

	return rs
}

func possibleValues(char byte) []string {
	if numbers[char] {
		return []string{string([]byte{char})}
	}

	c := string([]byte{char})
	return []string{strings.ToLower(c), strings.ToUpper(c)}
}
