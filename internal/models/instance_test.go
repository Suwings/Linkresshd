package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestInstanceJsonSave(t *testing.T) {
	a := &InstanceData{"AAA",  "java -jar xxxxx.jar nogui"}
	z, _ := json.Marshal(a)
	fmt.Println(string(z))

	f := z
	newObj := new(InstanceData)
	json.Unmarshal(f,newObj)

	fmt.Print("2:"+newObj.Name+" | " + " | "+ newObj.Command)

}
