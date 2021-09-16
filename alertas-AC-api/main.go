package main

import (
	"log"
	"os"
)

func main() {

	l := log.New(os.Stdout, "alertas-AC-api", log.LstdFlags)

	go func() {
		l.Println("Inciando servidor en ")
	}()
}
