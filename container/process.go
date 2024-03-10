package container

import (
	"os"
	"os/exec"
	"syscall"
)

// NewParentProcess 创建一个会隔离namespace进程的command
func NewParentProcess(tty bool) (*exec.Cmd, *os.File) {
	readPipe, writePipe, _ := os.Pipe()
	//调用资深，传入init参数，也就是执行initCommand
	cmd := exec.Command("/proc/self/exec", "init")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	cmd.ExtraFiles = []*os.File{
		readPipe,
	}
	return cmd, writePipe
}
