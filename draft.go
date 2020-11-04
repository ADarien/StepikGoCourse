package main

import (
	"strconv"
	"unicode"
)

func adding(s1, s2 string) int64 {
	string1 := ""
	for _, ch := range s1 {
		if unicode.IsDigit(ch) {
			string1 += string(ch)
		}
	}

	string2 := ""
	for _, ch := range s2 {
		if unicode.IsDigit(ch) {
			string2 += string(ch)
		}
	}
	ch1, _ := strconv.ParseInt(string1, 10, 64)
	ch2, _ := strconv.ParseInt(string2, 10, 64)
	res := ch1 + ch2
	// fmt.Println(res)

	return res
}

func main() {
	adding("%^80", "hhhhh20&&&&nd")
}
