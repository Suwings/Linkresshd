package mpty

import (
	"errors"
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/kr/pty"
	"io"
	"linkresshd/internal/service"
	"os"
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

var gProcess *service.MineProcess = nil
var gPtyFile *os.File = nil
var gSession *ssh.Session = nil

func GlobalExecProcess(pSession *ssh.Session, command string) (*service.MineProcess, error) {
	gSession = pSession
	ptyReq, _, isPty := (*gSession).Pty()
	if !isPty{
		return nil ,errors.New("not pty")
	}
	commandArr := strings.Fields(command)
	if len(commandArr) < 1 {
		return nil, errors.New("数组长度不正确，应该大于1实际却小于1")
	}
	if gProcess == nil {
		var err error
		gProcess := new(service.MineProcess)
		gProcess.Pcmd = exec.Command(commandArr[0], commandArr[1:]...)
		gProcess.Pcmd.Env = append(gProcess.Pcmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
		gPtyFile, err = pty.Start(gProcess.Pcmd)
		return gProcess, err
	}
	return gProcess, nil
}

func BindGlobalProcessIO(pSession *ssh.Session, mp *service.MineProcess) error {
	s := *gSession
	_, winCh, isPty := s.Pty()
	if isPty {
		go func() {
			for win := range winCh {
				SetWinsize(gPtyFile, win.Width, win.Height)
			}
		}()
		go func() {
			io.Copy(gPtyFile, s) // stdin loop
		}()
		io.Copy(s, gPtyFile) // stdout loop
	} else {
		return errors.New("No PTY requested")
	}
	return nil
}
