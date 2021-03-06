package main

import (
	"os"
	"fmt"
	"io"
)

func main(){
	file,err := os.Open("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	p := make([]byte,4)

	for {
		n,err := file.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
}