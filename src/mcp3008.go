package gopitools

import (
	"fmt"
	"os"
	"os/exec"
)

// Mcp3008 provides an interface to a MCP3008 chip, an 8 channel 10-bit ADC chip.
// This is currently a wrapper for a python routine until a pure golang solution
// can be found that provides similar functionality.
type Mcp3008 struct {
}

// Close releases any resources.
func (m *Mcp3008) Close() {
}

// Init initializes the device ready for use.
func (m *Mcp3008) Init() error {
	_, err := os.Stat("mcp3008.py")
	return err
}

// Read reads the current values from the ADC and returns them as a float slice
// containing 8 values.  Each value ranges from 0 to 1.
func (m *Mcp3008) Read() ([]float64, error) {
	if err := m.Init(); err != nil {
		return nil, err
	}
	if out, err := exec.Command("python", "mcp3008.py").Output(); err != nil {
		return nil, err
	} else {
		fmt.Println(string(out))
	}
	return nil, nil
}
