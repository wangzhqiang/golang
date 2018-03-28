package main

import "fmt"

type Pair struct{
	values [2]interface{}
}

func MakePair(k,v interface{}) Pair{
	return Pair{values:[2]interface{}{k,v}}
}

func main() {
	p := MakePair("Hello", false)

	fmt.Println(p.values[0],"",p.values[1])
}
