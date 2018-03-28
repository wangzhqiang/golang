package main

import (
	"bytes"
	"io"
	"strings"
	"os"
	"fmt"
	"compress/gzip"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int,error) {
	*c += ByteCounter(len(p))
	return len(p),nil
}

type GzipCounter int

func (g *GzipCounter) Write(p []byte) (int,error) {
	var buf bytes.Buffer
	zip := gzip.NewWriter(&buf)
	io.Copy(zip,strings.NewReader(string(p)))
	zip.Close()
	var c ByteCounter
	io.Copy(&c,&buf)
	*g += GzipCounter(c)
	return len(p),nil
}

func main() {
	var bc ByteCounter
	var gc GzipCounter
	tee := io.TeeReader(os.Stdin,&gc)
	io.Copy(&bc,tee)
	fmt.Println(float64(gc)/float64(bc))
}