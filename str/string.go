package str

import "strings"

func LowerFirstLetter(s string) string {
	s1 := strings.TrimSpace(s)
	if s1 == "" {
		return s
	}

	data := make([]byte, 0, len(s1))

	if s1[0] >= 'A' && s1[0] <= 'Z' {
		data = append(data, s1[0]+32)
		data = append(data, s1[1:]...)
		return string(data)
	}

	return s
}
