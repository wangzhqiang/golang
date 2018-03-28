package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	//2018-01-29 16:24:16.699490078 +0800 CST
	fmt.Printf("%02d.%02d.%4d\n",t.Day(),t.Month(),t.Year())
	//29.01.2018
	t = time.Now().UTC()
	fmt.Println(t)
	//2018-01-29 08:25:26.813383048 +0000 UTC
	fmt.Println(time.Now())
	//2018-01-29 16:25:54.706086652 +0800 CST
	c := time.Now().Unix()
	fmt.Println(c)
}
