package luhn

import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	sum := 0
	for j := 0; j < len(id); j++ {
		c := id[len(id)-1-j]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if j%2 == 1 {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}
	return sum%10 == 0
}