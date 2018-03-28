package main

import (
	"strings"
	"io"
	"os"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	i := io.TeeReader(r,os.Stdout)
	io.Copy(os.Stderr,i)
}
