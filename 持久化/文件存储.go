package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello World!\n")    //创建一个byte数组 有13个字符  [72 101 108 108 111 32 87 111 114 108 100 33 10]
	fmt.Println(data)
	err := ioutil.WriteFile("data1",data,0644)  //创建一个文件data1
	if err != nil{
		panic(err)
	}

	read1,_ := ioutil.ReadFile("data1")  //使用readFile方法读取文件
	fmt.Println(string(read1))


	data = []byte("Hello World!\n")
	file,_ := os.Create("data2")
	defer file.Close()
	bytes,_:= file.Write(data)
	fmt.Printf("Wrote %d bytes to file \n",bytes)

}
