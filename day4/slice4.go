package main

import "fmt"

func main() {
	var a [4]int = [4]int{2, 3, 4, 5}
	var b []int = a[1:3]
	fmt.Println(b)
}
