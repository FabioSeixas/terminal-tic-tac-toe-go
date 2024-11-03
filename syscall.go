package main

import (
	"syscall"
	"unsafe"
)

// https://www.delorie.com/djgpp/doc/libc/libc_495.html
// https://github.com/nsf/termbox-go/blob/master/termbox.go#L209

type winsize struct {
	rows    uint16
	cols    uint16
	xpixels uint16
	ypixels uint16
}

func get_term_size(fd uintptr) (int, int) {
	var sz winsize
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL,
		fd, uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))
	return int(sz.cols), int(sz.rows)
}
