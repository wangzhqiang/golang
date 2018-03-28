package main

import (
	"crypto/sha1"
	"io"
	"fmt"
)

func main(){
	hasher := sha1.New()
	io.WriteString(hasher,"test")
	b := []byte{}
	fmt.Printf("Result:%x\n",hasher.Sum(b))
	fmt.Printf("Result:%d\n",hasher.Sum(b))
}
