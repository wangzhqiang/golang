package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("c.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("hello world!")
	f.Close()
}
