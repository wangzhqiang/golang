package main

import (
	"os/exec"
	"fmt"
	"log"
)

func main() {
	cmd := exec.Command("ls","-l")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	s := string(stdoutStderr)
	fmt.Printf("%s\n", s)
}
