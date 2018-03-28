package main

import (
	"io"
	"errors"
	"os"
)

type LimitedReader struct{
	cnt int
	r   io.Reader
}

func (l *LimitedReader) Reader(p []byte)(int,error) {
	if l.cnt <= 0 {
		return 0,errors.New("EOF")
	}
	if l.cnt > 0 {
		p = p[:l.cnt]
	}
    return l.r.Read(p)
}

func NewLimitReader(r io.Reader,max int)(io.Reader) {
	nr := LimitedReader{max,r}
	return &nr
}

func main() {
	c := LimitedReader{5,os.Stdin}
	io.Copy(os.Stdout,NewLimitReader(c.r,5))
}
