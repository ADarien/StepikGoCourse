package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// str = strings.Trim(str, "\n")

	str := "1 149,6088607594936;1 179,0666666666666"
	lineName := strings.Split(string(str), ";")

	ch1 := strings.Replace(lineName[0], " ", "", 1)
	ch2 := strings.Replace(lineName[1], " ", "", 1)

	ch1 = strings.Replace(ch1, ",", ".", 1)
	ch2 = strings.Replace(ch2, ",", ".", 1)

	fmt.Println(ch1)
	fmt.Println(ch2)

	digit1, _ := strconv.ParseFloat(ch1, 64)
	digit2, _ := strconv.ParseFloat(ch2, 64)

	// fmt.Println(digit1 / digit2)
	res := digit1 / digit2

	text := strconv.FormatFloat(res, 'f', 4, 64)
	rs := []rune(text)

	fmt.Println(rs)

}
