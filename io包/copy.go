package main

import (
	"strings"
	"io"
	"os"
	"log"
)

func main(){
	r := strings.NewReader("some io.reader stream to be read\n")
	_,err := io.Copy(os.Stdout,r)
	if err != nil{
		log.Fatal(err)
	}
}
