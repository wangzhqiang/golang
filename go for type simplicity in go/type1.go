package main

import "fmt"

type Pair struct {
	values [2]interface{}
}

func MakePair(k,v interface{}) Pair {
	return Pair{values:[2]interface{}{k,v}}
}

func (p Pair) Get(i int) interface{} {
	return p.values[i]
}

func main() {
	p := MakePair("hello","flase")

	fmt.Println(p.Get(0),"",p.Get(1))
}