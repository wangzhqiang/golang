package main

import "fmt"

func main() {
	var a int
	a = 10
	switch a {
	case 1:
		fmt.Println(a)
	case 2:
		fmt.Println(a)
	default:
		fmt.Println("a","default",a)
	}
}
