package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

var procchan = make(chan *syscall.ProcAttr, 100)

func fork(spec *syscall.ProcAttr) {
	pid, err := syscall.ForkExec("./worker", []string{}, spec)
	if err != nil {
		fmt.Println(err)
	}

	p, _ := os.FindProcess(pid)
	fmt.Println(p)
	p.Wait()

	procchan <- spec
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8000")
	l, _ := net.ListenTCP("tcp4", addr)
	fd, _ := l.File()

	fds := []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd(), fd.Fd()}

	execSpec := &syscall.ProcAttr{
		Dir:   "/Users/snow/yan/lab/git/goutil",
		Env:   os.Environ(),
		Files: fds,
	}
	procchan <- execSpec

	for spec := range procchan {
		go fork(spec)
	}
}
