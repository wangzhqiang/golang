package main

import "fmt"

func main() {
	var a [6]int = [6]int{2, 3, 4, 5, 6, 7}
	var b []int = a[:4]
	fmt.Println("a的地址是\n")
	for i := 0; i < 6; i++ {
		fmt.Println(&a[i])
	}
	fmt.Println("b的地址是\n")
	for i := 0; i < 4; i++ {
		fmt.Println(&b[i])
	}
}
