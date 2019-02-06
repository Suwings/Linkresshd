package service

import (
	"encoding/json"
	"io/ioutil"
	"linkresshd/models"
	"log"
	"os"
	"path/filepath"
)

// Global user and command config
var GlobalConfigInstance models.ConfigInstance

func InitConfig() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if contents, err := ioutil.ReadFile(dir + "/config/config.json"); err == nil {
		log.Println("Load config:", dir+"/config/config.json")
		err := json.Unmarshal(contents, &GlobalConfigInstance)
		if err != nil {
			log.Println("Load config error:", err)
			panic(err)
		}
	} else {
		if err != nil {
			log.Println("Load config error:")
			panic(err)
		}
	}
	log.Println("[Config] Username:", GlobalConfigInstance.Name)
	log.Println("[Config] Command: ", GlobalConfigInstance.Command)
}
