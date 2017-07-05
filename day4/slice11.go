package main

import "fmt"

func main() {
	var a []int
	fmt.Printf("len=%d cap=%d a=%v\n", len(a), cap(a), a)
	if a == nil {
		fmt.Println("nil")
	}
}
