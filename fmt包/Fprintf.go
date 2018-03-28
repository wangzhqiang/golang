package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	fmt.Fprintf(os.Stdout,"%s\n","hello world! - unbuffered")

	buf := bufio.NewScanner(os.Stdout)

	fmt.Fprintf(buf,"%s\n","hello world! - buffered")

}
