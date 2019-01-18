package models

type InstanceData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Command  string `json:"command"`
}
