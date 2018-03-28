package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Account struct{
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Money     float64  `json:"money,string"`
}

var jsonString string = `{
	"email":"rsj217@gmail.com",
	"password":"123",
	"money":"100.5"
}`

func main() {
	account := Account{}

	err := json.Unmarshal([]byte(jsonString),&account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n",account)
}
