package main

import "fmt"

func main() {
	var a [6]int = [6]int{1, 4, 6, 3, 2, 1}
	var b []int = a[:4]
	fmt.Println(b)
}
