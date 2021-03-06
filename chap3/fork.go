package main

import (
	"fmt"
	"os"
	"syscall"
)

func child() {
	fmt.Println("I'm child! my pid is ", os.Getpid(), ".")
	os.Exit(0)
}

func parent(childPid uintptr) {
	fmt.Println("I'm parent! my pid is", os.Getpid(), "and the pid of my child is", childPid, ".")
	os.Exit(0)
}

func main() {
	ret1, ret2, err := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	fmt.Println("ret1", ret1)
	fmt.Println("ret2", ret2)
	fmt.Println("err", err)
	if err != 0 {
		fmt.Println("syscall.SYS_FORK failed")
	}
	if ret1 == 0 {
		// child process came here because fork() returns 0 for child process
		child()
	} else {
		// parent process came here because fork() returns the pid of newly created child process (>1)
		parent(ret1)
	}
	// should'nt reach here
	fmt.Println("shouldn't reach here")
}
