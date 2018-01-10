package main

import (
	"reflect"
	"fmt"
)

func main(){
	tag := reflect.StructTag(`species:"gopher" color:"blue"`)
	fmt.Println(tag.Get("color"),tag.Get("species"))
}
