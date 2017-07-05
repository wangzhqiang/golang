package main

import "fmt"

func main() {
	var a [4]string = [4]string{"hello", "wang", "zhi", "qiang"}
	fmt.Println(a)
	b := a[0:2]
	c := a[1:3]
	fmt.Println(b, c)
	fmt.Println(&a[0])

	b[0] = "zhao"
	fmt.Println(b, c)
	fmt.Println(a)
	fmt.Println(&a[0])
	fmt.Println(&b[0],&c[0])
}
