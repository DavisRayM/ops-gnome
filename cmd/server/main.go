package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/DavisRayM/ops-gnome/internal/ops"
	"github.com/DavisRayM/ops-gnome/pkg/config"
)

var configPath = flag.String("config", "ops.yaml", "Path to the config file.")

func main() {
	flag.Parse()

	c, err := config.LoadOpsConfig(*configPath)
	if err != nil {
		panic(err)
	}

	s, err := ops.New(c)
	if err != nil {
		panic(err)
	}

	errChan := make(chan error, 1)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	log.Println("Starting server at: ", c.Server.StringRep())
	go func() {
		defer close(errChan)
		errChan <- s.Start()
	}()

	select {
	case err := <-errChan:
		log.Println("Server exited with error: ", err)
	case <-stopChan:
		log.Println("Gracefully terminating server")
		s.Stop()
	}
}
