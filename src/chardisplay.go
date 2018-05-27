package gopitools

import (
	"os"
	"os/exec"
	"strings"
)

type CharDisplay struct {
	ClearOnClose bool
}

// Init initializes the Pin ready for use.
func (d *CharDisplay) Init() error {
	_, err := os.Stat("chardisplay.py")
	return err
}

// Clear clears the LCD display.
func (d *CharDisplay) Clear() error {
	if err := d.Init(); err != nil {
		return err
	}
	return exec.Command("python", "chardisplay.py", "-a", "clear").Run()
}

// Close releases and unmaps GPIO memory.
func (d *CharDisplay) Close() {
	if d.ClearOnClose {
		d.Clear()
	}
}

// Message writes the text to the display.
// A NewLine character '\n' in the text moves the rest of the text
// tn the next line.
func (d *CharDisplay) Message(text string) error {
	if err := d.Init(); err != nil {
		return err
	}
	lines := strings.Split(text, "\n")

	if len(lines) == 1 {
		return exec.Command("python", "chardisplay.py", "-l1", lines[0]).Run()
	} else {
		return exec.Command("python", "chardisplay.py", "-l1", lines[0], "-l2", lines[1]).Run()
	}
}
