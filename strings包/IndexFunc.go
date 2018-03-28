package main

import (
	"unicode"
	"fmt"
	"strings"
)

func main(){
	f := func(c rune) bool {
		return unicode.Is(unicode.Han,c)
	}

	fmt.Println(strings.IndexFunc("Hello, 中国",f))
	fmt.Println(strings.IndexFunc("Hello, world",f))
}