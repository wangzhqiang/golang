package main

import (
	"fmt"
	"strings"
)

func main() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return r
		case r >= 'a' && r <= 'z':
			return r - 32
		}
		return r
	}
	fmt.Println(strings.Map(rot13,"Twas brilling and the slithy gopher"))
}
