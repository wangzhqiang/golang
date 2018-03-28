package main

import "fmt"

func test(x int) func() {
	fmt.Println(&x)

	return func(){
		fmt.Println(&x,x)
	}
}

func main() {
	f := test(0x100)
	f()
}
