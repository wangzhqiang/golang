package main

import (
   "fmt"
   "strings"
   "flag"
)

var sep = flag.String("s","","separator")
func main() {
   flag.Parse()
   fmt.Print(strings.Join(flag.Args(),*sep))
    
}
