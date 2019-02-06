package service

import (
	"github.com/gliderlabs/ssh"
	"io"
	"linkresshd/mpty"
	"log"
)

func Authorization(s *ssh.Session) (bool, error) {
	username := (*s).User()
	_, err := io.WriteString(*s, "Password:")
	if err != nil {
		return false, err
	}
	password := mpty.ReadLine(s)
	if username != GlobalConfigInstance.Name || password != GlobalConfigInstance.Password {
		log.Println("Login failed. Username:", username, "address:", (*s).RemoteAddr().String())
		return false, nil
	} else {
		log.Println("Login successful. Username:", username, "address:", (*s).RemoteAddr().String())
		mpty.SendLine(s, "\n\n")
		return true, nil
	}
}
