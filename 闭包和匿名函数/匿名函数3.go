package main

import "fmt"

func test(f func()) {
	f()
}

func main() {
	test(func() {
		fmt.Println("hello world!")
	})
}