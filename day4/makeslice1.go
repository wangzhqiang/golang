package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice(a)
	a=append(a, 10)
	printSlice(a)
	a=append(a, 30)
	printSlice(a)
}

func printSlice(a []int) {
	fmt.Printf("len=%d cap=%d a=%v\n", len(a), cap(a), a)
}
