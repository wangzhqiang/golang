package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int,error) {
	*c += ByteCounter(len(p))
	return len(p),nil
}

func main() {
	var c ByteCounter
    fmt.Fprint(&c,"hello world")
    fmt.Println(c)
}