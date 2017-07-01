package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	f, err := os.Create("ls.out")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = f
	cmd.Start()
	cmd.Wait()
}
