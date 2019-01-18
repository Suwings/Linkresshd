package main

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"linkresshd/internal/controller"
	"linkresshd/internal/mpty"
	"log"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		username := s.User()
		//verification
		for !controller.Authorization(&s) {
			mpty.SendLine(&s, "\nUsername or password error, please try again.\n")
		}
		//welcome msg
		mpty.SendLine(&s, fmt.Sprintf("Welcome %s !\n", username))
		//exec system command
		//mpty.PtyExecProcess(&s,"/bin/bash")
		mp, err := mpty.GlobalExecProcess(&s, "/bin/bash")
		if err != nil {
			log.Print(err)
		}
		errw := mpty.BindGlobalProcessIO(&s, mp)
		if errw != nil {
			log.Print(err)
		}
	})

	log.Println("Starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
