package main

import "fmt"

func main() {
        m := map[string]int{
         "a":1,
         "b":2,
}
        fmt.Println(m["a"])
        fmt.Println(m["b"])
        delete(m,"a")
        fmt.Println(m["a"])
}


