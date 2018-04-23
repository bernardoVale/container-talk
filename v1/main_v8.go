package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("What?")
	}
}

func child() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Chroot("fs/springboot"))
	must(syscall.Mount("proc", "/proc", "proc", 0, ""))
	must(os.Chdir("/"))

	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
}

func parent() {
	fmt.Printf("running %v\n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
		Unshareflags: syscall.CLONE_NEWNS,
	}
	must(cmd.Run())
	//After that show that how we can use user ns to start the container with unprivilege user
}

// Exit with status code 1 in case of failure
func must(err error) {
	if err != nil {
		panic(err)
	}
}
