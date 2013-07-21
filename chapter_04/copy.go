package main

import (
	"syscall"
	"os"
	"fmt"
)

func main() {
	var inputFd, outputFd, openFlags int
	buf := make([]byte, 1024)
	if len(os.Args) != 3 {
		fmt.Println("usage")
		return
	}

	/* open input and output files */
	inputFd, err := syscall.Open(os.Args[1], syscall.O_RDONLY, 0)
	if inputFd == -1 || err != nil {
		fmt.Println("Error opening file : " +os.Args[1])
	}

	openFlags = syscall.O_CREAT | syscall.O_WRONLY | syscall.O_TRUNC
	filePerms := uint32(syscall.S_IRUSR | syscall.S_IWUSR | syscall.S_IRGRP | syscall.S_IWGRP | syscall.S_IROTH | syscall.S_IWOTH)

	outputFd, err = syscall.Open(os.Args[2], openFlags, filePerms)
	if inputFd == -1 || err != nil {
		fmt.Println("Error opening file : " +os.Args[2])
	}

	var numRead, numWritten int
	for {
		numRead, _ = syscall.Read(inputFd, buf)
		if numRead <= 0 {
			break
		}
		numWritten, err = syscall.Write(outputFd, buf)
		if numRead != numWritten || err != nil {
			fmt.Println("Could not write the entire buffer")
			return
		}
	}

	if numRead == -1 {
		fmt.Println("Error with read")
		return
	}
	
	if err = syscall.Close(inputFd); err != nil {
		fmt.Println("Error close input")
		return
	}

	if err = syscall.Close(outputFd); err != nil {
		fmt.Println("Error close output")
		return
	}
}
	
