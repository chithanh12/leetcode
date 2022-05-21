package reverse_string

import "strings"

func reverseWords(s string) string {
	tmp := make([]byte, 0)

	rs := []string{}
	for _, ss := range s {
		if ss == ' ' {
			reverseString(tmp)
			rs = append(rs, string(tmp))
			rs = append(rs, " ")
			tmp = make([]byte, 0)
		} else {
			tmp = append(tmp, byte(ss))
		}
	}

	if len(tmp) > 0 {
		reverseString(tmp)
		rs = append(rs, string(tmp))
	}
	return strings.Join(rs, "")
}
