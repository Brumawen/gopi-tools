package gopitools

import "testing"
import "time"

func TestCanInitialize(t *testing.T) {
	d := CharDisplay{}
	err := d.Init()
	if err != nil {
		t.Error(err)
	}
	defer d.Close()
}

func TestCanDisplayMessage(t *testing.T) {
	d := CharDisplay{}
	err := d.Message("Hello\nWorld")
	if err != nil {
		t.Error(err)
	}
	defer d.Close()
	time.Sleep(2 * time.Second)
	d.Clear()
}

func TestCanDisplayLongMessage(t *testing.T) {
	d := CharDisplay{}
	err := d.Message("This is a really long message.")
	if err != nil {
		t.Error(err)
	}
	defer d.Close()
	time.Sleep(2 * time.Second)
	d.Clear()
}

func TestCanDisplayRightToLeft(t *testing.T) {
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

func TestCanAutoScroll(t *testing.T) {
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

func TestCanShowCursor(t *testing.T) {
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
