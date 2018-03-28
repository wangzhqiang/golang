package main

import (
	"log"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {     //捕获异常
			log.Fatal(err)
		}
	}()
	panic("i am dead")            //引发错误
	fmt.Println("exit.")          //永不会执行
}
