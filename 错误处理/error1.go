package main

import (
	"errors"
	"log"
	"fmt"
)
//自己定义错误
var errDivByZero = errors.New("division by zero")

func div(x,y int) (int,error) {
	if y == 0 {
		return 0,errDivByZero
	}
	return x/y, nil
}

func main() {
	z, err := div(5,0)
	if err == errDivByZero {
		log.Fatal(err)
	}
	fmt.Println(z)
}
