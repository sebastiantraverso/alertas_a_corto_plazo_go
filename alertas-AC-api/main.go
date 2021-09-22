package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	config "github.com/sebastiantraverso/alertas_a_corto_plazo_go/alertas-AC-api/config"
)

func main() {

	l := log.New(os.Stdout, "alertas-AC-api", log.LstdFlags)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter() // TODO: crear func con los endpoints y que devuelva un mux

	// variables del server
	port, err := config.GetPort()
	if err != nil {
		l.Printf("Variables del server (port): Error - %s", err)
	}
	addSrv, err := config.GetServerAddress()
	if err != nil {
		l.Printf("Variables del server (addSrv): Error - %s", err)
	}

	// create a new server
	s := http.Server{
		Addr:         addSrv,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Inciando servidor en ", port)
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Iniciando server: Error iniciando - %s", err)
			os.Exit(0)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal: ", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
