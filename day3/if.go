package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	a = "abc123"
	n, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)

	if a == "1" {
		fmt.Println("yes")
	} else if a == "2" {
		fmt.Println("now")
	} else {
		fmt.Println(a)
	}
}
