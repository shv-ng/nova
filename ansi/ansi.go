package ansi

import "fmt"

const Escape = 27

var (
	ClearScreen = []byte{Escape, '[', 'H', Escape, '[', 'J'}

	EnableAlternativeBuffer  = "\033[?1049h"
	DisableAlternativeBuffer = "\033[?1049l"

	DisableReportingFocus = "\033[?1004l"
)

func MoveUp(step int) {
	fmt.Printf("\033[%d;A", step)
}
func MoveDown(step int) {
	fmt.Printf("\033[%d;B", step)
}
