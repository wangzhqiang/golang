package main

import (
	"os"
	"log"
	"fmt"
)

func main(){
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not GO\n",
		"Errors are values\n",
		"Don't panic",
	}

	file,err := os.Create("./proverbs.txt")    //创建文件
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _,p := range proverbs {
		n,err := file.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {             //判断写入是否成功
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println("file write done")
}
