package main

import "fmt"

type Book struct {
	Id   int
	name string
	bo   string
}

func main() {
	var a Book
	a.Id = 1
	a.name = "study"
	a.bo = "wang"
	fmt.Println(a)

	b := Book{
		Id:   2,
		name: "english",
		bo:   "qiang"}
	fmt.Println(b)
}
