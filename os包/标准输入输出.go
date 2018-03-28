package main

import (
	"os"
	"fmt"
)

/*
 var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)

func NewFile(fd uintptr, name string) *File

os.Stdout, os.Stdin, and os.Stderr, that are of type *os.File
*/

func main() {
	proverbs := []string{
		"Channels orchestrate mutexes serizlize\n",
		"Cgo is not go\n",
		"Errors are values\n",
		"Don't panic\n",
	}

	for _,p := range proverbs {
		n, err := os.Stdout.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
}