package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("f.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(f, "hello")
	fmt.Fprint(f, "hello\n")
	s := "i can do it"
	m := 5
	fmt.Fprintf(f, "s=%s,m=%d\n", s, m)
	f.Close()
}
