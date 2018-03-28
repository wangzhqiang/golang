package main

import "fmt"

type Pair struct{
	values [2]interface{}
}

func main(){
	p := Pair{values:[2]interface{}{"hello",false}}
	fmt.Println(p.values[0],"",p.values[1])
}


