package main

import "fmt"

func main() {
     s1 := "hello" + "world"
     s2 := "helloworld"
     if s1 == s2 {
      fmt.Println("equal")
}
     fmt.Println(0,len(s1)-1)
     var c byte
     c = s1[1]
     fmt.Println(c,s1,s2)
     s3 := s1[0:5]
     fmt.Println(s1,s2,s3)
     
}
