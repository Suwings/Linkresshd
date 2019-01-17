package controller

import (
	"linkresshd/internal/mpty"
	"fmt"
	"github.com/gliderlabs/ssh"
	"log"
)

func Authorization(s * ssh.Session) bool{
	username := (*s).User()
	mpty.SendLine(s, fmt.Sprintf("Username: %s , Password: ",username))
	password := mpty.ReadLine(s)
	if password != "foo" {
		log.Println("User failed. INFO:", username, password)
		return false
	}else {
		log.Println("User successful. INFO:", username, password)
		mpty.SendLine(s, "\n\n")
		return true
	}
}
