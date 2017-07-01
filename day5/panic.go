package main

import "fmt"

func main() {
        defer func() {
	 err:= recover()
         fmt.Println(err)
}()
	print()
}

func print() {
	var a *int
	fmt.Println(*a)
}
