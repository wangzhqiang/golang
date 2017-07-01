package main

import "fmt"

func main() {
	print()
}

func print() {
	var a *int
	fmt.Println(*a)
}
