package gopitools

import (
	"testing"
	"time"
)

func TestCharDisplayCanInitialize(t *testing.T) {
	d := CharDisplay{}
	err := d.Init()
	if err != nil {
		t.Error(err)
	}
	defer d.Close()
}

func TestCharDisplayCanDisplayMessage(t *testing.T) {
	d := CharDisplay{}
	err := d.Message("Hello\nWorld")
	if err != nil {
		t.Error(err)
	}
	defer d.Close()
	time.Sleep(2 * time.Second)
	d.Clear()
}

func TestCanDisplaySmallWords(t *testing.T) {
	d := CharDisplay{ClearOnClose: true}
	err := d.Message("I AM A\nA B C")
	if err != nil {
		t.Error(err)
	}
	defer d.Close()
	time.Sleep(2 * time.Second)
}
