package main

import (
	"fmt"
	"strings"
)

//ContainsAny reports whether any Unicode code points in chars are within s.

func main() {
	fmt.Println(strings.ContainsRune("aardvark",97))
	fmt.Println(strings.ContainsRune("timeout",97))
}