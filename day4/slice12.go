package main

import "fmt"

func main() {
	var a []int = []int{1}
	a = a[:0]
	fmt.Println(len(a))
	if a == nil {
		fmt.Println("nil")
	}
}
