package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("start")

	fd := os.NewFile(3, "")
	l, _ := net.FileListener(fd)
	l1 := l.(*net.TCPListener)

	for {
		c, err := l1.AcceptTCP()
		fmt.Println(c.RemoteAddr())
		if err != nil {
			fmt.Println(err)
		}
		go handle(c)
	}
}

func handle(c *net.TCPConn) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, 4)
	c.Write(append(buf, []byte("test")...))
	fmt.Println("ok")
	c.Close()
}
