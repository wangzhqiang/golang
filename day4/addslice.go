package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4}
	var b []int = make([]int, 6)
	a = append(a, b...)
	fmt.Println(a)
	fmt.Println(b)
}
