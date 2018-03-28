package main

import (
	"fmt"
	"strings"
)

//Contains reports whether substr is within s.

func main() {
	fmt.Println(strings.Contains("seafood","foo"))
	fmt.Println(strings.Contains("seafood","bar"))
	fmt.Println(strings.Contains("seafood",""))
	fmt.Println(strings.Contains("",""))
}