package main

import "fmt"

type pairVal [2]interface{}

type Pair struct{
	values pairVal
}

func main(){
	p := Pair{pairVal{"hello",false}}
	fmt.Println(p.values[0],"",p.values[1])
}
