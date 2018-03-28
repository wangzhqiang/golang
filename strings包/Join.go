package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"foo","bar","baz"}
	fmt.Println(strings.Join(s,","))
	fmt.Println(strings.Join(s,""))
}
