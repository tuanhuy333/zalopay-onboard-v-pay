package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"pay-service/cmd/server"
	"pay-service/pkg/config"
)

func main() {
	cp := os.Getenv("CONFIG_PATH")
	if cp == "" {
		log.Fatal("Config path is empty")
	}
	var c server.Config

	if err := config.Load(cp, &c); err != nil {
		log.Fatalf("Load config from %v failed: %v", cp, err)
	}

	s := server.NewServer(c)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, os.Interrupt)

	go s.Start()

	<-shutdown
	s.Shutdown()
}
