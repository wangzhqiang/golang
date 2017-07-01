package main

import "fmt"

func fib(a int) int {
	if a == 1 || a == 2 {
		return 1
	}
	return fib(a-1) + fib(a-2)

}

func main() {
	fmt.Print(fib(20))
}
