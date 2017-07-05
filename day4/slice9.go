package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 5, 4, 8, 6}
	printSlice(s)
	s = s[:0]
	printSlice(s)
	s = s[:4]
	printSlice(s)
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Println("len=%d cap=%d %v\n", len(s), cap(s), s)
}
