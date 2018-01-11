package main

import (
	"bytes"
	"fmt"
	"os"
)

func main(){
	proverbs := []string{
		"Channels orchestrate mutexes serailize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}


	var write bytes.Buffer

	for _,p := range proverbs {
		n, err := write.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println(write.String())
}
