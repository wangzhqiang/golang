package main

import "fmt"

func main() {
	c := make(map[string]int)
	c["a"] = 1
	c["b"] = 2
	fmt.Println(c)
	fmt.Println(c["a"])
	fmt.Println(c["b"])
	d, ok := c["c"]
	if ok {
		fmt.Println(d)
	} else {
		fmt.Println("not found")
	}
}
