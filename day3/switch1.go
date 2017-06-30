package main

import (
	"fmt"
	"os"
)

func main() {
	var n, m int
	n = 50
	m = 6
	switch os.Args[1] {
	case "+":
		fmt.Println(m + n)
	case "-":
		fmt.Println(m - n)
	case "*":
		fmt.Println(m * n)
	case "%":
		fmt.Println(m % n)
	}
}
