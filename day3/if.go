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

}
