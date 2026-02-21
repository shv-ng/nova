package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shv-ng/nova/terminal"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("fail to start", err)
	}
}

func run() error {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	t, err := terminal.New()
	if err != nil {
		return err
	}

	for range 15 {
		go star(t.Height, t.Width)
	}

	<-c

	return t.Restore()
}

func star(height, width int) {
	for {
		line := rand.Intn(height) + 1
		col := rand.Intn(width) + 1
		sleep := rand.Intn(5) + 2

		nova.PutText(".", line, col)
		time.Sleep(time.Second * time.Duration(sleep))
		nova.PutText(" ", line, col)
	}
}
