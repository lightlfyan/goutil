package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fd := os.NewFile(0, "")
	l, _ := net.FileListener(fd)
	l1 := l.(*net.TCPListener)

	for {
		c, err := l1.AcceptTCP()
		if err != nil {
			fmt.Println(c, err)
		}
	}
}
