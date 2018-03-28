package main

func worker(ch chan string) {
	for url := range ch {
		_ = url
	}
}

func main() {
	urls := []string{} //1000个url
	//以并发为10来抓取url
	ch := make(chan string)
	for i:=0;i<10;i++ {
		go worker(ch)
	}
	for _,url := range urls {
		ch <- url
	}
}