package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	b := a[0:3]
	c := a[0:3]
	d := c[0:2]
	e := b[0:2]
           
	fmt.Println(&a[0])
	fmt.Println(&b[0])
	fmt.Println(&c[0])
	fmt.Println(&d[0])
	fmt.Println(&e[0])
}
