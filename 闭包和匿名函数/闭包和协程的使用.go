package main

import (
	"fmt"
	"time"
	"net"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	for ix := range values {
		func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()

	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	time.Sleep(5e9)
}