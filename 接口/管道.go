package main

import (
	"strings"
	"io"
	"os"
	"bytes"
	"compress/gzip"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte)(int,error) {
	//string := string(p)
	s := strings.Split(string(p),",")
	return len(s),nil
}

type GzipCounter int

func (g *GzipCounter) Write(p []byte)(int,error) {
	var buf bytes.Buffer
	gzip
}

func main(){
	var c ByteCounter
	var buf bytes.Buffer

	i := io.TeeReader(os.Stdin,&buf)
	io.Copy(&c,i)

	gzReader,_ := gzip.NewReader(i)
	gzReader.Read()

}


