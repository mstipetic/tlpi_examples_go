package main

import (
	"fmt"
	"os"
	"syscall"
	"strconv"
)

func main() {
	fmt.Println("hi")
	if len(os.Args) < 3 {
		fmt.Println("Usage")
		return
	}

	fd, err := syscall.Open(os.Args[1], syscall.O_RDWR | syscall.O_CREAT ,syscall.S_IRUSR | syscall.S_IWUSR | syscall.S_IRGRP | syscall.S_IWGRP | syscall.S_IROTH | syscall.S_IWOTH)

	if fd == -1 || err != nil {
		fmt.Println("Error when opening file")
		fmt.Println(err)
		return
	}

	for ap := 2; ap < len(os.Args); ap++ {
		switch os.Args[ap][0] {
			case 'r', 'R': /* Display bytes at current offset, either as text or hex */
				length, _ := strconv.ParseInt(os.Args[ap][1:], 10, 64)
				buf := make([]byte, length)
				fmt.Println(len(buf))

				numRead, err := syscall.Read(fd, buf)
				if err != nil || numRead == -1 {
					fmt.Println("Error when reading from file")
					return
				}

				if numRead == 0 {
					fmt.Println("End of file")
				} else {
					fmt.Println(os.Args[ap])
					for i := 0; i < numRead; i++ {
						if os.Args[ap][0] == 'r' {
							fmt.Printf("%c", buf[i])
						} else {
							fmt.Printf("%02x ", buf[i])
						}
					}
					fmt.Println()
				}
			case 'w': /* write string at current offset */
				numWritten, err := syscall.Write(fd, []byte(os.Args[ap][1:]))
				if numWritten == -1 || err != nil {
					fmt.Println("Write error")
					return
				}
				fmt.Printf("%s: wrote %d bytes\n", os.Args[ap], numWritten)
			case 's':
				offset, _ := strconv.ParseInt(os.Args[ap][1:], 10, 64)
				fileOffset, err := syscall.Seek(fd, offset, os.SEEK_SET)
				if fileOffset == -1 || err != nil {
					fmt.Println("Seek error")
					return
				}
				fmt.Printf("%s: Seek succeded\n", os.Args[ap])
			default:
				fmt.Printf("Argument must start with [rRws]: %s\n", os.Args[ap])
		}
	}
	return
	
}
