// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "unicode"


// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var m map[string]struct{} = make(map[string]struct{})
    res := ""
    if len(s) > 1 {
        res += string(s[0])
    }
    for _, r := range s {
        if r == ' ' || r == '-' || r == '_'{
            m[" "] = struct{}{}
            continue
        }
        if _, ok := m[" "]; ok {
            res += string(unicode.ToUpper(r))
            delete(m, " ")
        }
        
    }
	return res
}
