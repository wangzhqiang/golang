package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for a, v := range a {
		fmt.Printf("%d %d\n", a, v)
	}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}
