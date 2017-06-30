package main

import "fmt"

func main() {
	var i int
	i = 10
	for {
		i = i + 1
		fmt.Println(i)
		if i > 20 {
			break
		}
	}
}
