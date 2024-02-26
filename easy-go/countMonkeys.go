package main

import "fmt"

func monkeyCount(n int) []int {
	result := make([]int, n)
	for i := 1; i <= n; i++ {
		result[i-1] = i
	}

	return result
}
func main() {
	result1 := monkeyCount(10)
	result2 := monkeyCount(1)

	fmt.Println(result1)
	fmt.Println(result2)
}
