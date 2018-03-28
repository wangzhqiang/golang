package main

import "fmt"

//直接执行

func main() {
	func(s string) {
		fmt.Println(s)
	}("hello, world!")
}
