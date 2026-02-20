package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shv-ng/nova/ansi"
	"golang.org/x/term"
)

func main() {

	if err := run(); err != nil {
		log.Fatalln("fail to start", err)
	}
}

func run() error {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	fd := os.Stdout.Fd()

	state, err := term.GetState(int(fd))
	if err != nil {
		return err
	}

	fmt.Println(ansi.EnableAlternativeBuffer)
	fmt.Println(ansi.DisableReportingFocus)

	fmt.Println("hello")
	time.Sleep(time.Second)
	ansi.MoveDown(5)
	fmt.Println("hi")

	<-c

	fmt.Println(ansi.DisableAlternativeBuffer)

	if err := term.Restore(int(fd), state); err != nil {
		return err
	}

	return nil

}
