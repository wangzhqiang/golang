package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	for ix := range values {
		val := values[ix]
		go func( ) {
			fmt.Print(val, " ")
		}()
	}
	time.Sleep(5e9)

}

