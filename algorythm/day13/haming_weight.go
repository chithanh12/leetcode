package day13

import "fmt"

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}
	return hammingWeight(num>>1) + int((num & 1))
}

func reverseBits(num uint32) uint32 {
	v := num << 31
	fmt.Printf("%b\n", v)
	fmt.Printf("%b\n", num>>31)
	for i := 1; i < 32; i++ {
		v = v | ((num >> i) << (31 - i))
	}

	return v
}
