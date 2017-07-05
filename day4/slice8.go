package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 5, 6, 7}
	s = s[1:5]
	fmt.Println(s)
	s = s[:3]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)
}
