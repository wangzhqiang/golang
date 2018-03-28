package main

import "fmt"

func test() func(int,int) int {
	return func(x,y int) int {
		return x + y
	}
}

func main() {
	add := test()
	fmt.Println(add(1,2))
}
