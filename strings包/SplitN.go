package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n",strings.SplitN("a,b,c,d,e,f",",",3))
	z := strings.SplitN("a,b,c",",",0)
	fmt.Printf("%q(nil = %v)\n",z,z == nil)
}