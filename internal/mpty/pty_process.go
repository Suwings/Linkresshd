package mpty

import (
	"errors"
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/kr/pty"
	"io"
	"os/exec"
	"strings"
)

func PtyExecProcess(pSession *ssh.Session, command string) (*exec.Cmd, error) {
	s := *pSession
	commandArr := strings.Fields(command)
	if len(commandArr) < 1 {
		return nil, errors.New("数组长度不正确，应该大于1实际却小于1")
	}
	cmd := exec.Command(commandArr[0], commandArr[1:]...)
	ptyReq, winCh, isPty := s.Pty()
	if isPty {
		cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
		f, err := pty.Start(cmd)
		if err != nil {
			return nil, err
		}
		go func() {
			for win := range winCh {
				SetWinsize(f, win.Width, win.Height)
			}
		}()
		go func() {
			io.Copy(f, s) // stdin loop
		}()
		io.Copy(s, f) // stdout loop
	} else {
		return nil, errors.New("No PTY requested")
	}
	return cmd, nil
}
