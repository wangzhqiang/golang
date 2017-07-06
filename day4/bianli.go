package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for name, age := range m {
		fmt.Println("name", name, "age", age)
	}
	for name := range m {
		fmt.Println(name)
	}
}
