package gopitools

import (
	"fmt"
	"testing"
)

func TestMcp3008CanInitialize(t *testing.T) {
	m := Mcp3008{}
	if err := m.Init(); err != nil {
		t.Error(err)
	}
	defer m.Close()
}

func TestMcp3008CanReadChannels(t *testing.T) {
	m := Mcp3008{}
	if err := m.Init(); err != nil {
		t.Error(err)
	}
	defer m.Close()

	r, err := m.Read(0, 1, 2)
	if err != nil {
		t.Error(err)
	}
	if len(r) == 0 {
		t.Error("No values returned")
	} else {
		fmt.Println(r)
	}
}
