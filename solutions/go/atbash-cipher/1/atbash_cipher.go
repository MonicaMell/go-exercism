package atbashcipher

import "strings"

func Atbash(s string) string {
	s = strings.ToLower(s)

	var cipher strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= 'a' && c <= 'z':
			cipher.WriteByte('a' + 'z' - c)
		case c >= '0' && c <= '9':
			cipher.WriteByte(c)
		}
	}

    t := cipher.String()
	var result strings.Builder
	for i := 0; i < len(t); i += 5 {
		if i > 0 {
			result.WriteByte(' ')
		}
		j := min(i+5, len(t))
		result.WriteString(t[i:j])
	}
	return result.String()
}
