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

	var s1 *Student
	s1 = &s
	s1.Id = 2
	fmt.Println(s)

	var s2 *int
	s2 = &s.Id
	*s2 = 5
	fmt.Println(s)
}
