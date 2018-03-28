package main

import "fmt"

//赋值给变量

func main() {
	add := func(x,y int) int {
		return x + y
	}
	fmt.Println(add(1,2))
}