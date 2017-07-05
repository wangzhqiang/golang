package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 4, 5, 6, 7}
	slice1 := a[1:3]
	fmt.Println(slice1)
}
