package main

import "fmt"

func main() {
	s1 := "hello" + "world"
	fmt.Println(s1)
	array := []byte(s1)
	fmt.Println(array)
	array[1] = 'k'
	s1 = string(array)
	fmt.Println(s1)
}
