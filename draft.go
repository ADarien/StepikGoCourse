package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swap func
func Swap(x []int, y int) {
	var a int
	a = x[y]
	x[y] = x[y+1]
	x[y+1] = a
}

// BubbleSort func
func BubbleSort(sl []int) {
	for i := range sl {
		for j := 0; j < len(sl)-i-1; j++ {
			if sl[j] > sl[j+1] {
				Swap(sl, j)
			}
		}
	}
}

func main() {
	fmt.Println("Input 10 numbers:")
	numbers := bufio.NewReader(os.Stdin)
	readStr, _, _ := numbers.ReadLine()
	nums := strings.Split(string(readStr), " ")

	var numSlice []int
	for _, num := range nums {
		i, _ := strconv.Atoi(num)
		numSlice = append(numSlice, i)
	}
	BubbleSort(numSlice)
	fmt.Println(numSlice)

}
