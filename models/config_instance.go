package models

type ConfigInstance struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Command  string `json:"command"`
	Port     int    `json:"port"`
}
