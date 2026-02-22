package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shv-ng/nova/config"
	"github.com/shv-ng/nova/terminal"
)

func main() {

	if err := run(); err != nil {
		log.Fatalln("fail with error", err)
	}
}

func run() error {

	countPtr := flag.Int("count", 100, "number of stars to display")
	durationPtr := flag.Int("duration", 4, "time (in s) to fade away a star")
	configPtr := flag.String("config", "", "load config file (*.py)")
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	t, err := terminal.New()
	if err != nil {
		return err
	}
	defer t.Restore()

	if err := terminal.DisableInputBuf(); err != nil {
		return err
	}

	if err := startStarts(*configPtr, *countPtr, *durationPtr, t.Height, t.Width); err != nil {
		return err
	}

	<-c

	return nil
}

func startStarts(configPtr string, countPtr, durationPtr, height, width int) error {
	if configPtr != "" {
		cfg, err := config.Load(configPtr)
		if err != nil {
			return err
		}
		if cfg.StarCount < 1 {
			cfg.StarCount = 100
		}
		if cfg.Symbol == "" {
			cfg.Symbol = "."
		}
		if cfg.FGColor == "" {
			cfg.FGColor = "255,255,255"
		}

		for idx := range cfg.StarCount {
			if idx < len(cfg.Stars) {
				s := cfg.Stars[idx]
				if s.Symbol == "" {
					s.Symbol = cfg.Symbol
				}
				if s.Duration < 1 {
					s.Duration = 4
				}
				if s.FGColor == "" {
					s.FGColor = cfg.FGColor
				}
				go star(s.Symbol, s.FGColor, height, width, s.Duration)
				continue
			}
			go star(cfg.Symbol, cfg.FGColor, height, width, cfg.Duration)
		}
		return nil
	}
	if countPtr < 1 {
		countPtr = 100
	}
	if durationPtr < 1 {
		durationPtr = 4
	}
	for range countPtr {
		go star(".", "255,255,255", height, width, durationPtr)
	}
	return nil
}
func star(symbol, fg string, height, width, duration int) {
	for {
		line := rand.Intn(height) + 1
		col := rand.Intn(width) + 1

		sleep := rand.Intn(3) + duration

		terminal.PutText(symbol, fg, line, col)
		time.Sleep(time.Second * time.Duration(sleep))
		terminal.PutText("  ", fg, line, col)
	}
}
