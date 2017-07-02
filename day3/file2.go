package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("b.txt")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("hello \n")
	f.Close()
}
