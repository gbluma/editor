// +build windows

package osutil

import (
	"errors"
	"os/exec"
)

//----------

const EscapeRune = '^'

//----------

func SetupExecCmdSysProcAttr(cmd *exec.Cmd) {
}

func KillExecCmd(cmd *exec.Cmd) error {
	return errors.New("todo: windows implementation")
}

//----------

func ShellRunArgs(args ...string) []string {
	return args
	//return []string{"bash", "-exec", strings.Join(args, " ")}
}

//----------

func GoExec() string {
	return ExecName("go")
}
func ExecName(name string) string {
	return name + ".exe"
}
