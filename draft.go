package main

import (
	"fmt"
	"unicode"
)

func main() {
	var text string
	fmt.Scan(&text)
	rs := []rune(text)
	var isLatin, isOkLetter, isLong bool
	for i := range rs {
		if unicode.Is(unicode.Latin, rs[i]) || unicode.IsLetter(rs[i]) || unicode.IsDigit(rs[i]) {
			isOkLetter = true
		} else {
			isOkLetter = false
			break
		}
		if unicode.Is(unicode.Latin, rs[i]) || unicode.IsDigit(rs[i]) {
			isLatin = true
		} else {
			isLatin = false
			break
		}

	}
	if len(rs) >= 5 {
		isLong = true
	}

	if isLatin && isOkLetter && isLong {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}

	fmt.Println(isLatin, isOkLetter, isLong)
}
