package main

import "fmt"

func print() {
	fmt.Println("hello")
}

func main() {
	var a func()
	a = print
	a()
}
