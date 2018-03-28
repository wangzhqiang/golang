package main

import (
	"os/exec"
	"log"
	"fmt"
)

func main(){
	cmd,err :=exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(cmd)
	fmt.Println(s)
}

