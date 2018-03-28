package main

import (
	"unicode"
	"fmt"
	"strings"
)


func main() {
	f := func(c rune) bool {
		return !unicode.IsNumber(c) && !unicode.IsLetter(c)
	}

	fmt.Println(strings.FieldsFunc("foo1;bar2,baz3...",f))
}
