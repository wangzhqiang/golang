package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

//读取文件
func main() {
	bytes,err := ioutil.ReadFile("./planers.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s",bytes)
}
