package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexByte("golang",'g'))
	fmt.Println(strings.IndexByte("gopher",'h'))
	fmt.Println(strings.IndexByte("golang",'x'))
}
