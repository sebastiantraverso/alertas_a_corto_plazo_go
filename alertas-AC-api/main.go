package main

import (
	"log"
	"os"

	config "github.com/sebastiantraverso/alertas_a_corto_plazo_go/alertas-AC-api/config"
)

func main() {

	l := log.New(os.Stdout, "alertas-AC-api", log.LstdFlags)

	go func() {
		l.Println("Inciando servidor en ", config.GetPort())

	}()
}
