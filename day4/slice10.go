package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 2, 6, 7, 8}
	a = a[:0]
	fmt.Printf("len=%d cap=%d a=%v\n", len(a), cap(a), a)
	a = a[5:6]
	fmt.Printf("len=%d cap=%d a=%v\n", len(a), cap(a), a)
}
