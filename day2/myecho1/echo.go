package main

import (
   "fmt"
   "strings"
   "flag"
)

var sep = flag.String("s","","separator")
var newline = flag.Bool("n",true,"append newline")
func main() {
   flag.Parse()
   fmt.Print(strings.Join(flag.Args(),*sep))
   if *newline {
   fmt.Println()
}
    
}
