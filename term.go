package nova

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/term"
)

type terminal struct {
	fd    int
	state *term.State

	Height, Width int
}

func NewTerminal() (*terminal, error) {
	fd := int(os.Stdout.Fd())

	state, err := term.GetState(fd)
	if err != nil {
		return nil, err
	}

	w, h, err := term.GetSize(fd)
	if err != nil {
		return nil, err
	}

	fmt.Print(EnableAltBuf)
	fmt.Print(DisableReportingFocus)
	fmt.Print(HideCursor)

	return &terminal{
		fd:     fd,
		state:  state,
		Height: h,
		Width:  w,
	}, nil
}

func (t *terminal) Restore() error {
	fmt.Print(DisableAltBuf)
	fmt.Print(ShowCursor)

	return term.Restore(t.fd, t.state)
}

func PutText(s string, line, col int) {
	MoveCursor(line, col)
	fmt.Printf("\033[5m%s\033[25", s)
}

func DisableInputBuf() error {
	return exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}
