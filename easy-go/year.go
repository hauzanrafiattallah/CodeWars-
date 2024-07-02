package main

import "fmt"

func century(year int) int {
	if year%10 == 0 {
		return year / 100
	} else {
		return year/100 + 1
	}
}

func main() {
	fmt.Println(century(1700))
}
