package main

import "fmt"

type File struct{
	fd    int
	name  string
}

func NewFile(fd int,name string) *File {
	if fd <0 {
		return nil
	}
	return &File{fd,name}
}

func main() {
	a := NewFile(1,"good")
	fmt.Print(*a)
}