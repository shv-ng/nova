package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/shv-ng/nova"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("fail to start", err)
	}
}

func run() error {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	t, err := nova.NewTerminal()
	if err != nil {
		return err
	}

	<-c

	return t.Restore()
}
