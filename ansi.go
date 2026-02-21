package nova

import (
	"fmt"
)

type seq string

const (
	EnableAltBuf  seq = "\033[?1049h"
	DisableAltBuf seq = "\033[?1049l"

	DisableReportingFocus seq = "\033[?1004l"

	ClearScreen seq = "\033[2J"
	MoveToHome  seq = "\033[H"

	HideCursor seq = "\033[?25l"
	ShowCursor seq = "\033[?25h"
)

func (s seq) Exec() {
	fmt.Println(s)
}

func Move(line, col int) {
	MoveToHome.Exec()
	fmt.Printf("\033[%d;%dH", line, col)
}
