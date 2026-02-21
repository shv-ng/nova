package cfg

import (
	_ "embed"
	"os/exec"
)

//go:embed extractor.py
var extractorScript string

func Load(file string) error {
	exec.Command("python", "-", extractorScript)
	return nil
}
