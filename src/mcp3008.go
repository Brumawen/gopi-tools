package gopitools

import (
	"fmt"
	"os"
	"os/exec"
)

type Mcp3008 struct {
}

func (m *Mcp3008) Close() {
}

func (m *Mcp3008) Init() error {
	_, err := os.Stat("mcp3008.py")
	return err
}

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
