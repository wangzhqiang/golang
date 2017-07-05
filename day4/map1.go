package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["a"] = 1
	ages["b"] = 2

	ages1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(ages)
	fmt.Println(ages1)
}
