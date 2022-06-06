package day11

import (
	"bytes"
	"fmt"
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
	rs := permutation([]byte(s))
	result := make([]string, 0, len(rs))
	fmt.Println("+++++++++++++++++++")

	for _, ss := range rs {
		result = append(result, string(ss))
		fmt.Println(ss)
	}

	return result
}

func permutation(inputs []byte) [][]byte {
	if len(inputs) == 1 {
		rs := make([][]byte, 0)
		for _, p := range possibleValues(inputs[0]) {
			rs = append(rs, []byte{p})
		}

		return rs
	}

	tmp := permutation(inputs[:len(inputs)-1])
	pos := possibleValues(inputs[len(inputs)-1])
	rs := make([][]byte, 0)

	fmt.Printf("pos =%v, val = %v \n", pos, string(inputs[len(inputs)-1]))

	for _, tt := range tmp {
		for _, p := range pos {
			xx := append(tt, p)
			rs = append(rs, xx)
		}
	}

	return rs
}

func possibleValues(char byte) []byte {
	if numbers[char] {
		return []byte{char}
	}

	return append(bytes.ToLower([]byte{char}), bytes.ToUpper([]byte{char})...)
}
