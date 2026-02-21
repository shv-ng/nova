package terminal

import (
	"fmt"
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
	fmt.Print(MoveToHome)
	fmt.Printf("\033[%d;%dH", line, col)
}
