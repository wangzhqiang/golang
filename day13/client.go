package main

import (
	"net"
	"os"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"io"
)

func main() {
	state, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatal(err)
	}
	defer terminal.Restore(0,state)

	conn,err := net.Dial("tcp",os.Args[1])
	if err != nil {
		fmt.Printf("%s",err)
		return
	}
	defer conn.Close()
	go io.Copy(conn,os.Stdin)
	io.Copy(os.Stdout,conn)

}