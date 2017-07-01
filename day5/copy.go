package main

import (
	"io"
	"log"
	"os"
)

func main() {
	v2()
}

func v2() {
	var f *os.File
	var err error
	if len(os.Args) > 1 {
		f, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	} else {
		f = os.Stdin
	}
	io.Copy(os.Stdout, f)
}
