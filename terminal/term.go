package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"golang.org/x/term"
)

type termi struct {
	fd    int
	state *term.State

	Height, Width int
}

func New() (*termi, error) {
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

	return &termi{
		fd:     fd,
		state:  state,
		Height: h,
		Width:  w,
	}, nil
}

func (t *termi) Restore() error {
	fmt.Print(DisableAltBuf)
	fmt.Print(ShowCursor)

	return term.Restore(t.fd, t.state)
}

var mu sync.Mutex

func PutText(s, fg string, line, col int) {
	mu.Lock()
	defer mu.Unlock()

	MoveCursor(line, col)
	ColorFg(fg)

	fmt.Printf("%s", s)
	ColorReset()

	os.Stdout.Sync()

}

func DisableInputBuf() error {
	return exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}
