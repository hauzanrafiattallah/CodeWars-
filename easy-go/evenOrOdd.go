package main

import "fmt"

func EvenOrOdd(num int) string {
	if num%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	fmt.Println(EvenOrOdd(2))

}
