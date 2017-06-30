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

	var n1 int32
	n1 = 10300000
	var n2 int8
	n2 = int8(n1)
	fmt.Println(n1, n2)
}
