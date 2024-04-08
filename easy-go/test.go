package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan nama anda: ")
	nama, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s\n", nama)
}
