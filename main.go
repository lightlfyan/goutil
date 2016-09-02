package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"time"
)

func fork(spec *syscall.ProcAttr) {
	pid, _ := syscall.ForkExec("dumpworker", []string{}, spec)
	p, _ := os.FindProcess(pid)
	fmt.Println(p)
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8000")
	l, _ := net.ListenTCP("tcp4", addr)
	fd, _ := l.File()
	fds := []uintptr{fd.Fd()}

	execSpec := &syscall.ProcAttr{
		Dir:   "/Users/snow/Downloads/testfd/",
		Env:   os.Environ(),
		Files: fds,
	}

	for i := 0; i < 2; i++ {
		go fork(execSpec)
	}

	time.Sleep(time.Second * 3)
}
