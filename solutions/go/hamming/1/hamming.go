package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) { return 0, errors.New("The length of two strings is not equal")}
	var d int = 0
    for i := 0; i < len(a) ; i++{
        if a[i] != b[i] {
            d++
        }
    }
    return d, nil
}
