package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	file, err := os.Create("/magic_msg.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = io.WriteString(file,"Go is fun")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

