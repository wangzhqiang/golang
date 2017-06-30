package main

import "fmt"

func main() {
	var x, y,s int64
	x = 1
	y = 1
        s =2
	for a := 0; a < 100; a++ {
		x = x + y
                s += x
		y = x + y
                s +=y
		fmt.Println(x)
		fmt.Println(y)
		fmt.Println(s)
		if y >= 100 {
			break
		}
	}
}
