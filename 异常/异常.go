package main

import (
	"runtime"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(1)
	for {
		int_chan := make(chan int,1)
		string_chan := make(chan string,1)
		int_chan <- 1
		string_chan <- "hehe"
		select {
		case value := <- int_chan:
			fmt.Println(value )
		case value := <- string_chan:
			panic(value + " out put")

		}
	}
}
