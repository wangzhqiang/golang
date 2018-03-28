package main

import "fmt"

func test(x int) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	f := test(123)
	f()
}