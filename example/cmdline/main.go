package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func system(command string) error {
	cmd := exec.Command("cmd")
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf(`/c "%s"`, command)}
	return cmd.Run()
}

func main() {
	if err := system(`notepad "foo.txt"`); err != nil {
		panic(err)
	}
}
