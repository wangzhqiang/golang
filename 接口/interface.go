package main

import (
	"strings"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) write(p []byte)(int,error) {
	//string := string(p)
	s := strings.Split(string(p),",")
	return len(s),nil
}

func main() {
	c := ByteCounter(0)
	fmt.Println(c.write([]byte("hello,world")))

}

