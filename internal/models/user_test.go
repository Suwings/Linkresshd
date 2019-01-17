package models

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	a := &User{"AAA",  []string{"A","B","CCC"}}
	z, _ := json.Marshal(a)
	fmt.Println(string(z))

	f := z
	newObj := new(User)
	json.Unmarshal(f,newObj)

	log.Print("2:"+newObj.Name+" | " + " | ", newObj.Collect)

}
