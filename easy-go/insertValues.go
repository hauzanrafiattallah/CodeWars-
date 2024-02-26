package main

import "fmt"

func Invert(arr []int) []int {
	inverted := make([]int, len(arr))
	for i, num := range arr {
		inverted[i] = -num
	}
	return inverted
}

func main() {
	input := []int{-1, 2, -3, 3, 2, -1, 0}
	fmt.Println(Invert(input))
}
