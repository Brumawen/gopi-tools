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

func TestCharDisplayCanDisplayLongMessage(t *testing.T) {
	d := CharDisplay{}
	err := d.Message("This is a really long message.")
	if err != nil {
		t.Error(err)
	}
	defer d.Close()
	time.Sleep(2 * time.Second)
	d.Clear()
}

func TestCharDisplayCanDisplayRightToLeft(t *testing.T) {
	d := CharDisplay{}
	err := d.SetRightToLeft()
	if err != nil {
		t.Error(err)
	}
	defer d.Close()
	d.Message("Hello\nWorld")
	time.Sleep(2 * time.Second)
	d.Clear()
}

func TestCharDisplayCanAutoScroll(t *testing.T) {
	d := CharDisplay{}
	err := d.AutoScroll(true)
	if err != nil {
		t.Error(err)
	}
	defer d.Close()
	d.Message("This is a really long message.")
	time.Sleep(2 * time.Second)
	d.Clear()
}

func TestCharDisplayCanShowCursor(t *testing.T) {
	d := CharDisplay{}
	err := d.ShowCursor(true)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(1 * time.Second)
	d.BlinkCursor(true)
	time.Sleep(1 * time.Second)
	d.BlinkCursor(false)
	time.Sleep(1 * time.Second)
	d.ShowCursor(false)
	d.Clear()
}
