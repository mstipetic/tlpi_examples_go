package main

import (
	"fmt"
	"syscall"
	"unsafe"
	"os"
)

type iovec struct {
	iov_base uintptr
	iov_len int
}

func main() {
/*
	openFlags := syscall.O_CREAT | syscall.O_WRONLY | syscall.O_TRUNC
	filePerms := uint32(syscall.S_IRUSR | syscall.S_IWUSR | syscall.S_IRGRP | syscall.S_IWGRP | syscall.S_IROTH | syscall.S_IWOTH)

	outputFd, err := syscall.Open(os.Args[1], openFlags, filePerms)
	if outputFd == -1 || err != nil {
		fmt.Println("Error opening file : " +os.Args[2])
	}

	buf := []byte{'t','e','s','t'}
	syscall.Syscall(syscall.SYS_WRITE, uintptr(outputFd), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	syscall.Close(outputFd)
	
*/
	if len(os.Args) != 2 {
		fmt.Println("usage")
		return
	}
	
	var fd int
	iov := make([]iovec, 3)

	buf1 := make([]byte, 6)
	var buf2 int
	buf3 := make([]byte, 100)

	fd, err := syscall.Open(os.Args[1], syscall.O_RDONLY, 0)
	if fd == -1 || err != nil {
		fmt.Println("Error open")
		return
	}

	totRequired := 0

	iov[0].iov_base = uintptr(unsafe.Pointer(&buf1[0]))
	iov[0].iov_len = len(buf1)
	totRequired += iov[0].iov_len

	iov[1].iov_base = uintptr(unsafe.Pointer(&buf2))
	iov[1].iov_len = 1
	totRequired += iov[1].iov_len
	
	iov[2].iov_base = uintptr(unsafe.Pointer(&buf3[0]))
	iov[2].iov_len = len(buf3)
	totRequired += iov[2].iov_len

	fmt.Println("before")
	fmt.Println(buf1)
	fmt.Println(buf2)
	fmt.Println(buf3)

	n, _, errn := syscall.Syscall(syscall.SYS_READV, uintptr(fd), uintptr(unsafe.Pointer(&iov[0])), uintptr(len(iov)))
	if errn != 0 {
		fmt.Println("Syscall error")
		fmt.Println(err)
		return
	}
	numRead := int(n)

	if numRead < totRequired {
		fmt.Println("Read fewer bytes than requested")
	}

	fmt.Println("after")
	fmt.Println(buf1)
	fmt.Println(buf2)
	fmt.Println(buf3)

	fmt.Printf("Total bytes requested : %d, bytes read : %d\n", totRequired, numRead)
	return
}
