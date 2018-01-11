package main

import (
	"strings"
	"io"
	"log"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte,33)

	if _,err := io.ReadAtLeast(r,buf,4); err != nil {
		log.Fatal(err)
	}

	shortBuf := make([]byte,3)
	if _,err := io.ReadAtLeast(r, shortBuf, 4);err != nil {
		log.Fatal(err)
	}




}
