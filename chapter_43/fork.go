package main

import (
	"fmt"
	"syscall"
)

func main() {
	s := make([]string, 0)
	fmt.Println("prije")
	pipes := make([]int, 2)
	ret := syscall.Pipe(pipes)
	fmt.Println(ret)
	fmt.Println(pipes)
	syscall.Dup2(1, pipes[0])
	syscall.ForkExec("/bin/ls", s, nil)
	buf := make([]byte, 200)
	syscall.Read(pipes[1], buf)
	fmt.Println("poslije")
	fmt.Println(buf)
}
