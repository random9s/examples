package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	n, err := whyAmIDoingThis()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Got", n)
}

func whyAmIDoingThis() (int, error) {
	//CREATE PIPE TO COMMUNICATE BETWEEN PROCESSES
	var p [2]int
	p[0] = -1
	p[1] = -1

	err := syscall.Pipe(p[:]) // for a good explanation of syscall pipe read this: https://www.geeksforgeeks.org/pipe-system-call/
	if err != nil {
		return 0, fmt.Errorf("pipe: %s", err)
	}
	syscall.CloseOnExec(p[0])
	syscall.CloseOnExec(p[1])

	//CREATE LS COMMAND
	var lsBin = "/bin/ls"
	var lsArgs = []string{"/bin/ls", "tmp"}
	lsStdin, _ := newNullReader()
	syscall.CloseOnExec(lsStdin)
	lsStderr, _ := newNullWriter()
	syscall.CloseOnExec(lsStderr)

	var lsProcAttr = &syscall.ProcAttr{
		Files: []uintptr{
			uintptr(lsStdin),
			uintptr(p[1]),
			uintptr(lsStderr),
		},
	}

	//CREATE WC COMMAND
	var wcBin = "/usr/bin/wc"
	var wcArgs = []string{"/usr/bin/wc", "-l"}
	wcStdout, _ := newStdout()
	syscall.CloseOnExec(wcStdout)
	wcStderr, _ := newNullWriter()
	syscall.CloseOnExec(wcStderr)

	var wcProcAttr = &syscall.ProcAttr{
		Files: []uintptr{
			uintptr(p[0]),
			uintptr(wcStdout),
			uintptr(wcStderr),
		},
	}

	lspid, err := syscall.ForkExec(lsBin, lsArgs, lsProcAttr)
	if err != nil {
		return 0, fmt.Errorf("fork exec: %s", err)
	}

	wcpid, err := syscall.ForkExec(wcBin, wcArgs, wcProcAttr)
	if err != nil {
		return 0, fmt.Errorf("fork exec: %s", err)
	}

	syscall.Wait4(lspid, nil, 0, nil)

	//READ FROM THE PIPE
	var b = make([]byte, 64)
	_, err = syscall.Read(wcStdout, b[:])
	if err != nil {
		return 0, fmt.Errorf("read: %s", err)
	}
	fmt.Println("read", string(b))

	syscall.Wait4(wcpid, nil, 0, nil)

	return 0, nil
}

func newStdin() (int, error) {
	return syscall.Open("/dev/stdin", syscall.O_RDONLY, 0)
}

func newStdout() (int, error) {
	return syscall.Open("/dev/stdout", syscall.O_WRONLY, 0)
}

func newStderr() (int, error) {
	return syscall.Open("/dev/stderr", syscall.O_WRONLY, 0)
}

func newNullReader() (int, error) {
	return syscall.Open("/dev/null", syscall.O_RDONLY, 0)
}

func newNullWriter() (int, error) {
	return syscall.Open("/dev/null", syscall.O_WRONLY, 0)
}

func sysCloseFd(fd int) {
	syscall.CloseOnExec(fd)
}
