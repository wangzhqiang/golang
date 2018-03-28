package main

import (
	"io"
	"log"
	"os"
)

func CountingWriter(w io.Writer)(io.Writer,*int64) {
	var c io.Writer
	var bytes []byte
	var d int64
	buf := make([]byte,1024)
	for {
		n, err := w.Write(buf)
		if err != nil {
			log.Fatal(err)
		}
		bytes = append(bytes,buf[:n]...)
	}
	n,err := c.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	d = int64(n)
	return c,&d
}

func main(){
	CountingWriter(os.Stdout)
}