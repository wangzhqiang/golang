package main

import (
	"strings"
	"bufio"
	"fmt"
)

//打印有多少字符串

type ByteCounter int

func (c *ByteCounter) Write(p []byte)(int, error) {
	buf := strings.NewReader(string(p))
	s := bufio.NewScanner(buf)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		fmt.Println(s.Text())
		*c++
	}
	return int(*c),nil
}

func main() {
	var b ByteCounter
	fmt.Fprint(&b,"hello it is a good thing")
	fmt.Println(b)
}