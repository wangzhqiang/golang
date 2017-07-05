package main

import "fmt"

func main() {
	var a [10]int = [10]int{1, 3, 4, 6, 3, 2, 1, 8, 9, 34}
	for k, v := range a {
		fmt.Printf("%d %d\n", k, v)
	}
}
