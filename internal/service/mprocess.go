package service

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

type MineProcess struct {
	Pstdin  *io.WriteCloser
	Pstdout *io.ReadCloser
	Pstderr *io.ReadCloser
	Pcmd    *exec.Cmd
}

func NewMineProcess(command string, params []string) (*MineProcess ,error){
	mp := new(MineProcess)
	return mp,mp.RunProcess(command, params)
}

func (p *MineProcess) RunProcess(commandName string, params []string) error {
	cmd := exec.Command(commandName, params...)
	p.Pcmd = cmd
	stdout, err1 := cmd.StdoutPipe()
	stdin, err2 := cmd.StdinPipe()
	stderr, err3 := cmd.StderrPipe()
	p.Pstdin = &stdin
	p.Pstdout = &stdout
	p.Pstderr = &stderr
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1,err2,err3)
		return err1
	}
	cmd.Start()
	return nil
}

func (p *MineProcess) StdoutLoop(ch chan string, callback func(string), readCloser *io.ReadCloser) {
	if readCloser == nil {
		readCloser = p.Pstdout
	}
	reader := bufio.NewReader(*readCloser)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			fmt.Println("Stdio Loop End")
			break
		}
		if callback != nil {
			callback(line)
		} else {
			ch <- line
		}
	}
}

func (p *MineProcess) ExecCommand(command string) {
	(*p.Pstdin).Write([]byte(command + "\n"))
}
