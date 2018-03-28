package main

import (
	"fmt"
	"strings"
)

//compare返回一个整数  a==b 返回0   a<b 返回-1 a>b 返回 1

func main() {
	fmt.Println(strings.Compare("a","b"))  //-1
	fmt.Println(strings.Compare("a","a"))  // 0
	fmt.Println(strings.Compare("b","a"))  // 1
}

