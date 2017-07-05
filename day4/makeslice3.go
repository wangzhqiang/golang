package main

import "fmt"

func main() {
	a := make([]int, 7)
	printSlice(a)
	b := make([]int, 6, 10)
	printSlice(b)
}

func printSlice(a []int) {
	fmt.Println("len=%d cap=%d a=%v", len(a), cap(a), a)
}
