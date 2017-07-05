package main

import "fmt"

func main() {
	var a [10]int = [10]int{2, 3, 4}
	for _, k := range a {
		fmt.Println(k)
	}
}
