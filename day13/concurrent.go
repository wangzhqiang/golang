package main

import (
	"net/http"
	"io"
	"os"
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func main() {
	urls := []string{} //1000个url
	//已并发为10来抓取url

	var url chan string
	go func(){
		for _,v:=range urls {
		wg.Add(1)
		url <- v
	  }
	}()


	for i:=0;i<10;i++ {
		go handle(url)
	}
	wg.Wait()
}

func handle(url chan string ){
	for {
		resp, err := http.Get(<-url)
		if err != nil {
			fmt.Println(err)
		}
		io.Copy(os.Stdout, resp.Body)
		wg.Done()
	}
}
