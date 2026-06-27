package isogram

import "unicode"

func IsIsogram(word string) bool {
	seen := map[rune]struct{}{}
    for _, r := range word {
        if unicode.IsLetter(r) {
            r = unicode.ToLower(r)
            if _, exists := seen[r] ; exists {
            	return false
        	}
       		seen[r] = struct{}{}
        }
    }
    return true
}
