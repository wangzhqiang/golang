package main

import "log"

func main() {
	defer func() {
		for {
			if err := recover(); err != nil {
				log.Fatal(err)
			} else {
				log.Fatal("fatal")
			}
		}
	}()

	defer func() {
		panic("you are dead")
	}()

	panic("i am dead")
}
