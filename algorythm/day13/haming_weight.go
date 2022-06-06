package day13

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}
	return hammingWeight(num>>1) + int((num & 1))
}
