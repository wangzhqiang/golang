package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	a = 10
	var s string
	s = strconv.Itoa(a)
	n, err := strconv.Atoi("123abc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Println(s)
}
