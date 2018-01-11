package main

import (
	"strings"
	"io"
	"os"
	"log"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")
	_,err := io.CopyN(os.Stdout,r,5)
	if err != nil {
		log.Fatal(err)
	}

}
