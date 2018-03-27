package gopitools

import "testing"
import "fmt"

func TestCanReadTemp(t *testing.T) {
	o := OneWireTemp{ID: "28-0516a4c75bff"}
	err := o.Init()
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
	o := OneWireTemp{ID: "28-0516a4c75bff"}
	err := o.Init()
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
