package command

import (
	"os"
	"path"
	"syscall"
)

func ExecCmd(cmd []string) error {
	if len(cmd) == 0 {
		return nil
	}
	name := path.Base(cmd[0])
	newArgs := append([]string{name}, cmd[1:]...)
	return syscall.Exec(cmd[0], newArgs, os.Environ())
}
