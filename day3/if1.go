package main

import (
	"fmt"
	"os"
        "strconv"
)

func main() {
	if os.Args[2] == "+" {
		fmt.Println(os.Args[1])
		fmt.Println(os.Args[2])
		fmt.Println(os.Args[3])
		fmt.Println(strconv.itoA(os.Args[1]) + strconv.itoA(os.Args[3]))
	}
}
