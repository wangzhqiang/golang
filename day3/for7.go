package main

import "fmt"

func main() {
	var s string
	s = "hello"
	for _, arr := range s {
		fmt.Println(arr)
	}
}
