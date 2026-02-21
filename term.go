package nova

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/shv-ng/nova"
	"golang.org/x/term"
)

type terminal struct {
	fd    int
	state *term.State

	height, width int
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

	EnableAltBuf.Exec()
	DisableReportingFocus.Exec()
	HideCursor.Exec()

	return &terminal{
		fd:     fd,
		state:  state,
		height: h,
		width:  w,
	}, nil
}

func (t *terminal) Restore() error {
	nova.DisableAltBuf.Exec()
	nova.ShowCursor.Exec()

	return term.Restore(t.fd, t.state)
}

func PutText(s string, line, col int) {
	MoveCursor(line, col)
	fmt.Println(s)
}

func DisableInputBuf() error {
	return exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}
