package main

import "fmt"

func swap(a string, b string) (string, string) {
	return b, a
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
