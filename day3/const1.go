package main

import "fmt"

func main() {
	const (
		PI = 3.1415926
		s  = 3.0
		m  = 3.5
	)
	const (
		A = iota
		B
		C
	)
	fmt.Println(A, B, C)
}
