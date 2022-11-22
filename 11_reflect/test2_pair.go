package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// sst: pair<type: *File, value: "/dev/tty描述符">
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open failed error", err)
		return
	}

	var r io.Reader
	// r: pair <type: *File, value: "/dev/tty描述符">
	r = tty

	var w io.Writer
	// w: pair <type: *File, value: "/dev/tty描述符">
	w = r.(io.Writer)

	w.Write([]byte("hello this is a tty!\n"))
}
