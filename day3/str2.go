package main

import "fmt"

func main() {
	//相加
	s := "hello" + "world"
	fmt.Println(s)
	//取字符
	var c byte
	c = s[2]
	fmt.Println(c)
	//切片
	s2 := s[0:5]
	fmt.Println(s2)

}
