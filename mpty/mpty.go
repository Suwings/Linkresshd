package mpty

import (
	"bufio"
	"io"
	"os"
	"syscall"
	"unsafe"

	"github.com/gliderlabs/ssh"
)

func SetWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

func SendLine(s *ssh.Session, str string) {
	io.WriteString(*s, str)
}

func ReadLine(s *ssh.Session) string {
	buf := bufio.NewReader(*s)
	resBytes := make([]byte, 2048)
	i := 0
	for {
		b, err := buf.ReadByte()
		if err != nil || b == '\n' || b == '\r' {
			break
		} else if b == '\b' || b == 127 {
			// delete a byte
			resBytes[i] = 0
			if i <= 0 {
				continue
			} else {
				i--
			}
		} else {
			// append a byte
			if i >= 2048 {
				break
			}
			resBytes[i] = b
			i++
		}
	}
	return string(resBytes[:i])
}
