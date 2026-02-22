package terminal

import (
	"fmt"
	"strings"
)

const (
	EnableAltBuf  = "\033[?1049h"
	DisableAltBuf = "\033[?1049l"

	DisableReportingFocus = "\033[?1004l"

	ClearScreen = "\033[2J"
	MoveToHome  = "\033[H"

	HideCursor = "\033[?25l"
	ShowCursor = "\033[?25h"
)

func MoveCursor(line, col int) {
	fmt.Printf("\033[%d;%dH", line, col)
}

func ColorFg(fg string) {
	if fg == "" {
		return
	}
	new := strings.ReplaceAll(fg, ",", ";")
	fmt.Printf("\033[38;2;%sm", new)
}

func ColorReset() {
	fmt.Print("\033[0m")
}
