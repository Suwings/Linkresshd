package main

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"linkresshd/mpty"
	"linkresshd/service"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	log.Println("Process pwd:", dir)
	//init
	service.InitConfig()
	//SSH Server handle
	ssh.Handle(func(s ssh.Session) {
		username := s.User()
		//verification
		for {
			isLogin, err := service.Authorization(&s)
			if err != nil {
				log.Println(" login error:", err)
				return
			}
			if isLogin {
				break
			}
			mpty.SendLine(&s, "\nUsername or password error, please try again.\n")
		}
		//welcome msg
		mpty.SendLine(&s, fmt.Sprintf("Welcome, %s.\n", username))
		//exec system command
		mpty.PtyExecProcess(&s, service.GlobalConfigInstance.Command)

	})
	log.Println("Starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
