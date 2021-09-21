package config

import "fmt"

const port int = 5000
const server string = "localhost"

//func init() {
//}

func GetPort() int {
	return port
}

func GetServerAddress() string {
	addr := fmt.Sprintf("%s:%d", server, GetPort())
	return addr
}
