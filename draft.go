package main

import "fmt"

func main() {
	a, b := sumInt(1, 4, 9, 6, 0)
	fmt.Println(a, b)
}

func sumInt(a ...int) (int, int) {
	var sum int
	for _, elem := range a {
		sum += elem
	}
	return len(a), sum
}
