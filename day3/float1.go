package main

import "fmt"

func main() {
	var a int
	var b float32
	a = 10
	b = float32(a / 3)
	fmt.Println(b * 3)
	a = int(b * 3)
	fmt.Println(b, a)
}
