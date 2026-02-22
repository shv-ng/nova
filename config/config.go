package config

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"os/exec"
)

//go:embed extractor.py
var extractorScript []byte

type cfg struct {
	StarCount int    `json:"star_count"`
	Symbol    string `json:"symbol"`
	Duration  int    `json:"duration"`

	BGColor string `json:"bg_color"`
	FGColor string `json:"fg_color"`

	Stars []Star `json:"stars"`
}

type Star struct {
	Symbol   string `json:"symbol"`
	Duration int    `json:"duration"`

	FGColor string `json:"fg_color"`
}

func Load(file string) (*cfg, error) {
	cmd := exec.Command("python", "-", file)
	cmd.Stdin = bytes.NewReader(extractorScript)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	c := cfg{}
	if err := json.Unmarshal(out, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
