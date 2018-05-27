package gopitools

import (
	"errors"
	"fmt"
	"testing"
)

func TestCanReadTemp(t *testing.T) {
	st, err := GetDeviceList()
	if err != nil {
		t.Error(err)
	}
	if len(st) == 0 {
		t.Error(errors.New("Temperature device not found."))
	} else {
		fmt.Println("Devices found", st)
	}

	o := OneWireTemp{ID: st[0].ID}
	err = o.Init()
	if err != nil {
		t.Error(err)
	}
	defer o.Close()
	v, err := o.ReadTemp()
	if err != nil {
		t.Error(err)
	}
	if v == 999999 {
		t.Error("Temperature could not be read from the file.")
	} else {
		fmt.Println("Temperature", v)
	}
}

func TestCanReadTempInFahrenheit(t *testing.T) {
	st, err := GetDeviceList()
	if err != nil {
		t.Error(err)
	}
	if len(st) == 0 {
		t.Error(errors.New("Temperature device not found."))
	}

	o := OneWireTemp{ID: st[0].ID}
	err = o.Init()
	if err != nil {
		t.Error(err)
	}
	defer o.Close()
	v, err := o.ReadTempInFahrenheit()
	if err != nil {
		t.Error(err)
	}
	if v == 999999 {
		t.Error("Temperature could not be read from the file.")
	} else {
		fmt.Println("Temperature", v)
	}
}

func TestCanGetDeviceList(t *testing.T) {
	lst, err := GetDeviceList()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(lst)
}
