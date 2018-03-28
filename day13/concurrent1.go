package main


func main() {
	urls := []string{}
	token := make(chan bool,10)
	for _,url := range urls {
		token <- true
		go func(url string) {
			_ = url
			<- token
		}(url)
	}
}
