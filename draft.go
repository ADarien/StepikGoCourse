package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	pos := 1
	file, err := os.Open("./task.data")
	if err != nil {
		fmt.Println(err)
	}

	br := bufio.NewReader(file)
	for num, err := br.ReadString(';'); err == nil; num, err = br.ReadString(';') {
		if num == "0;" || num == "0" {
			break
		}
		pos++
	}

	fmt.Println(pos)
}
