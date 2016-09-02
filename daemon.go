package main

import (
	"fmt"
	"os/exec"
    "flag"
)
var cmd *string = flag.String("cmd", "ls", "cmd")

func spawn(exePath string, counter int) error {
	fmt.Println("start ", exePath, counter)
	cmd := exec.Command(exePath, fmt.Sprintf("%d", counter))
	err := cmd.Start()
	if err != nil {
		return err
	}
	cmd.Wait()
	return nil
}

func main() {
    flag.Parse()
    fmt.Println(*cmd)
    return

	counter := 0
    exePath := *cmd
	#exePath := "./bin/master 1>master.log 2>&1"
	for {
		err := spawn(exePath, counter)
		if err != nil {
			fmt.Printf("spawn error: %s\n", err)
			return
		}
		counter += 1
	}
}
