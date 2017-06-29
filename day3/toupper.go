package main

import "fmt"

func main() {
	a := toupper("HelloWorld")
	fmt.Println(a)
}

func toupper(a string) string {
	var b int = len(a) 
	array := []byte(a)
	for i := 0; i < b; i++ {
		if array[i] > 96 && array[i] < 123 {
			array[i] = (array[i] - ('h'-'H'))
		}
	}
	a = string(array)
	return a
}
