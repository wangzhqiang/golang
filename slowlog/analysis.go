package main

import (
	"os"
	"log"
	"bufio"
	"strings"
)

func Readfile(slowlog string) {
	logfile,err := os.Open(slowlog)
	if err != nil {
		log.Fatal(err)
	}
    defer logfile.Close()
    Reader := bufio.NewReader(logfile)
	sqllog,err := os.Create("/Users/knowbox/go/src/github.com/wangzhqiang/golang/slowlog/sql.log")
	defer sqllog.Close()
   for {
	   line, err := Reader.ReadBytes('\n')
	   if err != nil {
		   log.Fatal(err)
	   }
	   if strings.HasPrefix(string(line),"select") {
	   	sqllog.Write(line)
	   }
   }
}

func main() {
	Readfile("/Users/knowbox/go/src/github.com/wangzhqiang/golang/slowlog/slow.log")
}
