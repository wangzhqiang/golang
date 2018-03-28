package main

import (
	"fmt"
	"strings"
)

//Count counts the number of non-overlapping instances of substr in s.
//If substr is an empty string, Count returns 1 + the number of Unicode code points in s.


func main() {
	fmt.Println(strings.Count("cheese","e"))
	fmt.Println(strings.Count("five",""))
}