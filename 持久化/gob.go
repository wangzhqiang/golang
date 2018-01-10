package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func store(data interface{},filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename,buffer.Bytes(),0600)
	if err != nil{
		log.Fatal(err)
	}
}

func load(data interface{}, filename string) {
	raw,err := ioutil.ReadFile(filename)
	if err != nil{
		log.Fatal(err)
	}
	buffer := bytes.NewReader(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		log.Fatal(err)
	}
}

func main(){
	post := Post{Id:1,Content:"Hello World!",Author:"Vanyarpy"}
	store(post,"post3")
	var postRead Post
	load(&postRead,"post3")
	fmt.Println(postRead)
}