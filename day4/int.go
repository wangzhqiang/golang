package main

import "fmt"

func main() {
	var a, b int
	a = 1
	b = a
	fmt.Println(&a, &b)
}
