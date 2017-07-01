package main

import "fmt"

func recu(n int) int {
	if n == 1 {
		return 2
	}
	return 2*recu(n-1) + n - 1
}

func main() {

	fmt.Println(recu(4))
}
