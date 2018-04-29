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

	r, err := m.Read()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
}
