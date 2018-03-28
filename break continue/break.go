package main

import "fmt"

func main() {
	var i int = 10
	for {
		i = i -1
		fmt.Printf("The variables i is now: %d\n",i)
		if i <0 {
			break
		}
	}
	fmt.Println("this is a bad")

}
