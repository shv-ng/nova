package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/shv-ng/nova/ansi"
	"golang.org/x/term"
)

func main() {

	if err := run(); err != nil {
		log.Fatalln("fail to start", err)
	}
}

func run() error {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	fd := os.Stdout.Fd()

	state, err := term.GetState(int(fd))
	if err != nil {
		return err
	}

	ansi.EnableAltBuf.Exec()
	ansi.DisableReportingFocus.Exec()
	ansi.HideCursor.Exec()

	ansi.Put("hi", 1, 1)
	ansi.Put("hello", 5, 15)

	<-c

	ansi.DisableAltBuf.Exec()
	ansi.ShowCursor.Exec()

	if err := term.Restore(int(fd), state); err != nil {
		return err
	}

	return nil

}
