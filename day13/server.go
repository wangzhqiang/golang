package main

import (
	"net"
	"log"
	"os/exec"
	"github.com/kr/pty"
	"io"
)

func handle(conn net.Conn) {
	defer conn.Close()
	//r:=bufio.NewReader(conn)
	//cmdstr,_ := r.ReadString('\n')
	cmd :=exec.Command("bash")
	fd,err := pty.Start(cmd)
	if err != nil {
		log.Fatal(err)
	}
	go io.Copy(fd,conn)
	io.Copy(conn,fd)

}

func main() {
	l,err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn,_ := l.Accept()
		go handle(conn)
	}
}
