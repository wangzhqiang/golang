package main

import "fmt"

func main() {
   str1 := "hello"
   str2 := "hello\n"
   str3 := "hello\tworld"
   str4 := "hello\"world\a"
   fmt.Println(str1)
   fmt.Println(str2)
   fmt.Println(str3)
   fmt.Println(str4)
}
