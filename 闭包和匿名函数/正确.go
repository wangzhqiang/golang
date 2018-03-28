package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	var names = [2]string{"1","2"}
	for _,name := range names {
		wg.Add(1)
		go func(n string) {
			fmt.Println(n)
			wg.Done()
		}(name)
	}
	wg.Wait()
}