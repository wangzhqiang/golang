package main

import "fmt"

func main() {
	var (
		a int
		b int32
		c int64
		d uint
		e uint32
		f uint64
	)
	fmt.Println("int默认值:", a)
	fmt.Println("int32默认值:", b)
	fmt.Println("int64默认值:", c)
	fmt.Println("uint默认值:", d)
	fmt.Println("uint32默认值:", e)
	fmt.Println("uint64默认值:", f)
}
