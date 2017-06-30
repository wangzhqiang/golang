package main

import "fmt"

func main() {
	var a, b int
	a = 10
	b = 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a = a + 3
	if b >= a {
		fmt.Println("hello world")
	} else {
		fmt.Println("yes")
	}
	if (a > 0 && b > 2) && a > 5 {
		fmt.Println("yes")
	}
	fmt.Println(a)
	c := a << 2
	fmt.Println(c)
}
