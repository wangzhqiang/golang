package main

import "fmt"
import "unsafe"

func main() {
	var (
		a int
		b int32
		c int64
		d uint
		e uint32
		f uint64
	)
        //查看整数的默认组
	fmt.Println("int默认值:", a)
	fmt.Println("int32默认值:", b)
	fmt.Println("int64默认值:", c)
	fmt.Println("uint默认值:", d)
	fmt.Println("uint32默认值:", e)
	fmt.Println("uint64默认值:", f)
        
        //查看整数的最大值
        //查看整数的字节数
        fmt.Println("int的字节数是:",unsafe.Sizeof(a))
        fmt.Println("int32的字节数是:",unsafe.Sizeof(b))
        fmt.Println("int64的字节数是:",unsafe.Sizeof(c))
        fmt.Println("uint的字节数是:",unsafe.Sizeof(d))
        fmt.Println("uint32的字节数是:",unsafe.Sizeof(e))
        fmt.Println("uint64的字节数是:",unsafe.Sizeof(f))
}
