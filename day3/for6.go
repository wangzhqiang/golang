package main

import "fmt"

func main() {
	var s string
	s = "hello"
	for i, arr := range s {
		fmt.Println(i, arr)
	}
}
