package main

import "fmt"

func add(n int) func(int) int {
      return func(m int) int {
         return m+n
}
}

func main() {
   f:=add(3)
   fmt.Println(f(2))
}
