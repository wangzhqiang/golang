package main

import (
	"strings"
	"io"
	"os"
	"log"
)

func main(){
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte,8)

	//buf被用
	_,err :=io.CopyBuffer(os.Stdout,r1,buf)
	if err != nil {
		log.Fatal(err)
	}
    //buf被重用  不用重新申请额外的buffer
	_,err = io.CopyBuffer(os.Stdout,r2,buf)
	if err != nil {
		log.Fatal(err)
	}
}
