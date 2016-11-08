package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
)

func keeprun(cmd *exec.Cmd) {
START:
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cmd.Process)
	cmd.Wait()

	dumpcmd := exec.Command(cmd.Path, cmd.Args...)
	dumpcmd.Stdout = os.Stdout
	dumpcmd.Stderr = os.Stderr
	dumpcmd.ExtraFiles = cmd.ExtraFiles

	cmd = dumpcmd
	goto START

}

func run(fd *os.File) {

	path := "/Users/snow/Downloads/afclient.py"
	cmd := exec.Command(path, ">>./log")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{fd}
	keeprun(cmd)
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8000")
	netListener, _ := net.ListenTCP("tcp4", addr)
	fd, _ := netListener.File()
	go run(fd)
	run(fd)
}
