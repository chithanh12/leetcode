package reverse_string

func reverseString(s []byte) {
	l := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[l-i] = s[l-i], s[i]
	}
}
