package main

import "fmt"

type Student struct {
	Id   int
	name string
}

func main() {
	var s Student
	s.Id = 1
	s.name = "wang"
	fmt.Println(s)

	s2 := Student{
		Id:   2,
		name: "qiang",
	}
	fmt.Println(s2)
}
