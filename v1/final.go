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
		run()
	case "child":
		child()
	default:
		panic("What ur doing?")
	}
}

func run() {
	fmt.Println("Forking a child process")
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// CLONE_NEWPID = PID Namespaces = Process tree isolation
	// CLONE_NEWNS = Mount Namespace
	// https://www.toptal.com/linux/separation-anxiety-isolating-your-system-with-linux-namespaces
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	must(cmd.Run())

}

func child() {
	command := os.Args[2]
	fmt.Printf("running %v as PID %d\n", command, os.Getpid())
	// fmt.Println("running", command)

	cmd := exec.Command(command, os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
	must(os.MkdirAll("rootfs/oldrootfs", 0700))
	must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
	must(os.Chdir("/"))

	// http://man7.org/linux/man-pages/man2/mount.2.html
	must(syscall.Mount("proc", "/proc", "proc", 0, ""))

	// must(syscall.Chroot("/vagrant/springboot/rootfs"))
	// must(os.Chdir("/"))
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
